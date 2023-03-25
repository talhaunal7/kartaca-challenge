package service

import (
	"example.com/auction-api/entity"
	"example.com/auction-api/model"
)

type ProductService interface {
	Add(add *model.ProductAdd) error
	GetById(id int) (*entity.Product, error)
	Offer(offer *model.ProductOffer, userId int) error
}
