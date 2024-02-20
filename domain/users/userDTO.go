package users

import (
	"strings"

	"anggi.tabulation/utils/errors"
)

type User struct {
	ID        int64  `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
