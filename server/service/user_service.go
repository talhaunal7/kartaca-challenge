package service

import (
	"example.com/auction-api/model"
)

type UserService interface {
	Register(*model.UserRegister) error
	Login(*model.UserLogin) (*model.UserDto, *string, error)
	Logout(id float64) error
}
