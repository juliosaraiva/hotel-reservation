package models

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 10
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

func (params UserParams) Validate() map[string]string {
	err := map[string]string{}
	if len(params.FirstName) < minFirstNameLen {
		err["first_name"] = fmt.Sprintf("first_name length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		err["last_name"] = fmt.Sprintf("last_name length should be at least %d characters", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		err["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	if !isEmailValid(params.Email) {
		err["email"] = "e-mail is invalid"
	}

	if len(err) > 0 {
		return err
	} else {
		return nil
	}
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func NewUserFromParams(params UserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}

func UpdateUserFromParams(params UserParams) (*UserUpdate, error) {
	var user UserUpdate
	if len(params.Password) > 0 {
		encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
		if err != nil {
			return nil, err
		}
		user.EncryptedPassword = string(encpw)
	}

	user.FirstName = params.FirstName
	user.LastName = params.LastName
	user.Email = params.Email

	return &user, nil
}
