package service

import "example.com/auction-api/model"

type ProductService interface {
	Add(add *model.ProductAdd) error
}
