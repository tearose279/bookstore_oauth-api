package services

import (
	"projectPath/domain/users"
	"projectPath/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr)  {
	result := &users.User{Id: userId}
	if err:= result.Get(); err != nil {
		return nil, err
	}
	return nil, nil
}
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateUser(); err != nil {
	return nil, err
	}
	if err := user.Save();err != nil {
		return nil,err
	}
	return &user, nil
}