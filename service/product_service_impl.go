package service

import (
	"errors"
	"example.com/auction-api/entity"
	"example.com/auction-api/model"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		db: db,
	}
}

func (prd *ProductServiceImpl) Add(productAddRequest *model.ProductAdd) error {

	product := entity.Product{
		Name:       productAddRequest.Name,
		OfferPrice: productAddRequest.OfferPrice,
		UserID:     30,
	}
	result := prd.db.Create(&product)
	if result.Error != nil {
		return errors.New("failed to add product")
	}
	return nil
}
