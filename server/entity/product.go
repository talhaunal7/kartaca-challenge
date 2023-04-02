package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	OfferPrice uint
	UserID     uint
	ImgUrl     string
	User       User `gorm:"foreignkey:UserID"` 
}
