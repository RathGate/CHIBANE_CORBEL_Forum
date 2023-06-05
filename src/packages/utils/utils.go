package utils

import (
	"github.com/asaskevich/govalidator"
)

func IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}
