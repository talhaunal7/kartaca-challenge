package service

import (
	"example.com/auction-api/model"
)

type ProductService interface {
	Add(add *model.ProductAdd) error
	GetAll() ([]*model.ProductDto, error)
	Offer(offer *model.ProductOffer, userId int) error
}
