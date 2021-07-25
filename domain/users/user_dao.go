package users

import (
	"fmt"
	"projectPath/utils/errors"
)

var (
	userDB = make(map[int64] *User)
)

func something()  {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
}

func (user User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func  (user *User)Save () *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewNotFoundError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already existed", user.Id))
	}
	userDB[user.Id] = user

	return nil
}