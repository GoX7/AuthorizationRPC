package validator

import (
	"errors"
	"net/mail"
)

func Email(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}
