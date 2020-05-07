package users

import (
	"fmt"
	"github.com/DevAgani/bookstore_users-api/utils/errors"
	"strings"
)

const (
	ActiveStatus = "active"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
	Status	string `json:"status"`
	Password string `json:"password"`
}

func (user *User) Validate() *errors.RestErr  {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == ""{
		return errors.NewBadRequestError("invalid email address")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == ""{
		return errors.NewBadRequestError("invalid password")
	}
	isPasswordValid := notValidatePassword(user.Password)
	if !isPasswordValid{
		return errors.NewBadRequestError(fmt.Sprintf("password should be more than 5 characters, you provided %d",len(user.Password)))
	}

	return nil
}

func notValidatePassword(pass string) bool{
	if len(pass) < 5 {
		return false
	}
	return true
}

