package service

import (
	"example.com/auction-api/entity"
	"example.com/auction-api/model"
)

type UserService interface {
	Register(*model.UserRegister) error
	Login(*model.UserLogin) (*entity.User, *string, error)
	Logout(string) error
	Validate() error
}
