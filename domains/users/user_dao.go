package users

import (
	"fmt"

	"github.com/culunvb/go_userapi/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NotFound(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	currentUser := usersDB[user.Id]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
