package model

type UserLogout struct {
	UserId string `json:"id" binding:"required"`
}
