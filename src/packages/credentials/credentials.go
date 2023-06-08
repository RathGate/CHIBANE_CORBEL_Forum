package credentials

import (
	"database/sql"
	"fmt"
	"forum/packages/dbData"
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

func ValidateUser(username, password string) (*dbData.User, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user dbData.User
	err = db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}
