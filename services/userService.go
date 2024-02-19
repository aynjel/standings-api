package services

import (
	"anggi.tabulation/domain/users"
	"anggi.tabulation/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to hash password")
	}

	user.Password = string(pwSlice[:])
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmailOrUsername(user users.User) (*users.User, *errors.RestErr) {
	if err := user.FindByEmailOrUsername(); err != nil {
		return nil, err
	}

	return &user, nil
}
