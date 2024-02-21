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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	user.Password = string(hashedPassword)

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

func GetAllUsers() ([]users.User, *errors.RestErr) {
	user := &users.User{}
	users, err := user.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByUsernameAndEmail(user users.User) (*users.User, *errors.RestErr) {
	if err := user.GetByUsernameAndEmail(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Update(currentUser); err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(userID int64) *errors.RestErr {
	user := &users.User{ID: userID}
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}
