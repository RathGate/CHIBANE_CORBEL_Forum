package credentials

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"unicode"

	"github.com/asaskevich/govalidator"
)

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
	regex := regexp.MustCompile(`^(?=.*[a-zA-Z])[a-zA-Z0-9_-]{3,20}$`)
	return regex.MatchString(username)
}

// TODO
// func LoginUser(username, password string) (user data.ShortUser, err error) {
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
// 	if err != nil {
// 		return user, err
// 	}
// 	defer db.Close()

// 	var tempID int
// 	var tempUsername string
// 	err = db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&tempID, &tempUsername)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("User not found")
// 			return user, err
// 		}
// 		fmt.Println(err)
// 		return user, err
// 	}

// 	if user.Password != password {
// 		return nil, fmt.Errorf("invalid password")
// 	}
// 	return &user, nil
// }

type FormField struct {
	Name     string `json:"name"`
	ErrorMsg string `json:"errorMsg"`
}
type FormValidation struct {
	Status        int         `json:"status"`
	InvalidFields []FormField `json:"fields"`
}

func CheckUserCredentials(username, password string) (formValidation FormValidation, userID int) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		formValidation.Status = http.StatusInternalServerError
		return formValidation, -1
	}
	defer db.Close()

	var tempID sql.NullInt64
	row := db.QueryRow(fmt.Sprintf(`SELECT id FROM users WHERE username = "%s" AND password = "%s"`, username, password))

	if err := row.Scan(&tempID); err != nil {
		if err == sql.ErrNoRows {
			formValidation.Status = http.StatusBadRequest
			formValidation.InvalidFields = append(formValidation.InvalidFields, FormField{
				Name:     "username",
				ErrorMsg: "Incorrect username or password",
			})
			return formValidation, -1
		}
		formValidation.Status = http.StatusInternalServerError
		return formValidation, -1
	}

	formValidation.Status = http.StatusOK
	return formValidation, int(tempID.Int64)
}

func RegisterNewUser(username, password, email string) (formValidation FormValidation, lastInserted int) {

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

	usernameExists, emailExists, errStatus := userExistsInDatabase(username, email)

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

	formValidation.Status, lastInserted = addUserToDatabase(username, password, email)

	return formValidation, lastInserted
}

func addUserToDatabase(username, password, email string) (status int, id int) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
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
	result, err = stmt.Exec(username, password, email, 3)
	if err != nil {
		return http.StatusInternalServerError, 0
	}

	lid, err := result.LastInsertId()
	if err != nil {
		return http.StatusInternalServerError, 0
	}

	return http.StatusOK, int(lid)
}

func userExistsInDatabase(username, email string) (usernameResult, emailResult bool, status int) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
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
