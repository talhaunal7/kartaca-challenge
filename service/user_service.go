package service

import (
	"example.com/auction-api/model"
)

type UserService interface {
	Register(*model.UserRegister) error
	Login(*model.UserLogin) (*string, error)
	Logout(string) error
	Validate() error
}
