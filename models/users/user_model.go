package users

import (
	"errors"
	"main/validators"
	userValidator "main/validators/users"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *User) Validate() error {
	if !userValidator.IsValidFirstName(u.FirstName) {
		return errors.New("invalid first name")
	} else if !userValidator.IsValidLastName(u.LastName) {
		return errors.New("invalid last name")
	} else if !validators.IsValidEmail(u.Email) {
		return errors.New("invalid email")
	} else if len(u.Password) < 8 {
		return errors.New("invalid password")
	} else if len(u.Password) > 255 {
		return errors.New("invalid password")
	}
	return nil
}