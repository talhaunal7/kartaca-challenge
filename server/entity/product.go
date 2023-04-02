package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	OfferPrice uint
	UserID     *int `sql:"DEFAULT:NULL"`
	ImgUrl     string
	User       User `gorm:"foreignkey:UserID"`
}
