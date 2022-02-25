package users

import (
	"strings"

	"github.com/culunvb/go_userapi/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	DateCreated string `json:"datecreated"`
}

// function structure
// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.BadRequestError("Invalid Email addresss")
// 	}
// 	return nil
// }

// method Structure
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Invalid Email addresss")
	}
	return nil
}
