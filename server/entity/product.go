package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	OfferPrice uint
	UserID     uint
	User       User `gorm:"foreignkey:UserID"`
}
