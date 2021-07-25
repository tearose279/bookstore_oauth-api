package users

import (
	"projectPath/utils/errors"
	"strings"
)

type User struct {
	Id int64
	FirstName string
	LastName string
	Email string
	DateCreated string
}

func ValidateUser(user *User) *errors.RestErr {
	user.Email =  strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}

func (user *User) ValidateUser() *errors.RestErr {
	user.Email =  strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}