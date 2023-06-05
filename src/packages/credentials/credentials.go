package credentials

import (
	"database/sql"
	"fmt"
	"forum/packages/users"
	"unicode"
)

func ContainsLetter(password string) bool {
	for _, char := range password {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func ContainsDigit(password string) bool {
	for _, char := range password {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func ContainsSpecialChar(password string) bool {
	for _, char := range password {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func ValidateUser(username, password string) (*users.User, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/forum?parseTime=true")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user users.User
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
