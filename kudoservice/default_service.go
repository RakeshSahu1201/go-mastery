package kudoservice

import (
	"errors"
	"main/kudotypes"
	"net/mail"
	"strings"
)

func ValidateUserRequest(u *kudotypes.UserRequest) []kudotypes.FieldError {
	if u == nil {
		return []kudotypes.FieldError{
			{Field: "request", Message: "user request cannot be nil"},
		}
	}

	var errors []kudotypes.FieldError
	if strings.TrimSpace(u.Name) == "" {
		errors = append(errors, kudotypes.FieldError{
			Field:   "name",
			Message: "name is required",
		})
	}

	if strings.TrimSpace(u.Email) == "" {
		errors = append(errors, kudotypes.FieldError{
			Field:   "email",
			Message: "email is required",
		})
	}

	return nil
}

func IsValidEmail(email string) error {
	email = strings.TrimSpace(email)
	if email == "" {
		return errors.New("invalid email: email can not be empty")
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return err
	}

	return nil
}
