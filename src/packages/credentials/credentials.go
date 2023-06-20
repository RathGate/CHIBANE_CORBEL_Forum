package credentials

import (
	"database/sql"
	"fmt"
	"forum/packages/utils"
	"net/http"
	"regexp"
	"unicode"

	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

// *FORM STRUCTS

type FormValidation struct {
	Status        int         `json:"status"`
	InvalidFields []FormField `json:"fields"`
}
type FormField struct {
	Name     string `json:"name"`
	ErrorMsg string `json:"errorMsg"`
}

// *FORMAT CHECKING FUNCTIONS

func ContainsAny(password string, f func(rune) bool) bool {
	for _, char := range password {
		if f(char) {
			return true
		}
	}
	return false
}

func ContainsLetter(password string) bool {
	return ContainsAny(password, unicode.IsLetter)
}

func ContainsDigit(password string) bool {
	return ContainsAny(password, unicode.IsDigit)
}

func ContainsSpecialChar(password string) bool {
	return ContainsAny(password, func(char rune) bool {
		return !unicode.IsLetter(char) && !unicode.IsDigit(char)
	})
}

func IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func IsValidUsername(username string) bool {
	regex := regexp.MustCompile(`^(.*[a-zA-Z])[a-zA-Z0-9_-]{3,20}$`)
	return regex.MatchString(username)
}

// *HASH FUNCTIONS

// Returns the hashed version of parameter password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Checks if string password corresponds to hashed password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// *LOGIN CHECKS AND FUNCTIONS

// Checks user credentials validity.
// Returns a json summary of the request, intended to be read by JS, as well as
// the ID of the user if existing.
func CheckUserCredentials(dba utils.DB_Access, username, password string) (formValidation FormValidation, userID int) {
	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		formValidation.Status = http.StatusInternalServerError
		return formValidation, -1
	}
	defer db.Close()

	var hashPassword string
	row := db.QueryRow(fmt.Sprintf(`SELECT id, password FROM users WHERE username = "%s"`, username))

	err = row.Scan(&userID, &hashPassword)

	if err == sql.ErrNoRows || !CheckPasswordHash(password, hashPassword) {
		formValidation.Status = http.StatusBadRequest
		formValidation.InvalidFields = append(formValidation.InvalidFields, FormField{
			Name:     "username",
			ErrorMsg: "Incorrect username or password",
		})
		return formValidation, -1
	} else if err != nil {
		formValidation.Status = http.StatusInternalServerError
		return formValidation, -1
	}
	formValidation.Status = http.StatusOK
	return formValidation, userID
}

// *REGISTER CHECKS AND FUNCTIONS

// Global function checking credentials and adding new user to DB if possible.
// Returns a json summary of the request (intended to be read by javascript),
// and if user successfully inserted, returns their ID as well.
func RegisterNewUser(dba utils.DB_Access, username, password, email string) (formValidation FormValidation, lastInserted int) {
	if len(password) < 8 {
		formValidation.InvalidFields = append(formValidation.InvalidFields, FormField{
			Name:     "password",
			ErrorMsg: "Password must be at least 8 characters long",
		})

	} else if !ContainsLetter(password) || !ContainsDigit(password) || !ContainsSpecialChar(password) {
		formValidation.InvalidFields = append(formValidation.InvalidFields, FormField{
			Name:     "password",
			ErrorMsg: "Password must contain at least one letter, one digit, and one special character",
		})
	}

	if !IsValidEmail(email) {
		formValidation.InvalidFields = append(formValidation.InvalidFields, FormField{
			Name:     "email",
			ErrorMsg: "Email address is not valid",
		})
	}

	usernameExists, emailExists, errStatus := userExistsInDatabase(dba, username, email)

	if errStatus != 200 {
		formValidation.Status = http.StatusBadRequest
		return formValidation, -1
	}
	if usernameExists {
		formValidation.InvalidFields = append(formValidation.InvalidFields, FormField{
			Name:     "username",
			ErrorMsg: "A user with this username already exists",
		})
	}
	if emailExists {
		formValidation.InvalidFields = append(formValidation.InvalidFields, FormField{
			Name:     "email",
			ErrorMsg: "A user with this email address already exists",
		})
	}

	if len(formValidation.InvalidFields) > 0 {
		formValidation.Status = http.StatusBadRequest
		return formValidation, -1
	}

	formValidation.Status, lastInserted = addUserToDatabase(dba, username, password, email, 4)

	return formValidation, lastInserted
}

// Adds user to database.
// Returns response status and newly-inserted user ID if successful
func addUserToDatabase(dba utils.DB_Access, username, password, email string, role_id int) (status int, id int) {
	hashPassword, err := HashPassword(password)
	if err != nil {
		return http.StatusInternalServerError, 0
	}

	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return http.StatusInternalServerError, 0
	}
	defer db.Close()

	var stmt *sql.Stmt
	stmt, err = db.Prepare("INSERT INTO users (username, password, email, role_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		return http.StatusInternalServerError, 0
	}
	defer stmt.Close()

	var result sql.Result

	result, err = stmt.Exec(username, hashPassword, email, role_id)

	if err != nil {
		return http.StatusInternalServerError, 0
	}

	lid, err := result.LastInsertId()
	if err != nil {
		return http.StatusInternalServerError, 0
	}
	return http.StatusOK, int(lid)
}

// Checks if password and email are already existing in the database
func userExistsInDatabase(dba utils.DB_Access, username, email string) (usernameResult, emailResult bool, status int) {
	db, err := sql.Open("mysql", dba.ToString())
	if err != nil {
		return false, false, http.StatusInternalServerError
	}
	defer db.Close()

	var tempID_1 sql.NullInt64
	row := db.QueryRow(fmt.Sprintf(`SELECT id FROM users WHERE username = "%s"`, username))

	if err := row.Scan(&tempID_1); err != nil && err != sql.ErrNoRows {
		return false, false, http.StatusInternalServerError
	}

	var tempID_2 sql.NullInt64
	row = db.QueryRow(fmt.Sprintf(`SELECT id FROM users WHERE email = "%s"`, email))

	if err := row.Scan(&tempID_2); err != nil && err != sql.ErrNoRows {
		return false, false, http.StatusInternalServerError
	}
	return tempID_1.Valid, tempID_2.Valid, http.StatusOK
}
