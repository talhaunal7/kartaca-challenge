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

func (prd *ProductServiceImpl) Add(productAddReq *model.ProductAdd) error {

	product := entity.Product{
		Name:       productAddReq.Name,
		OfferPrice: productAddReq.OfferPrice,
		UserID:     40,
	}
	result := prd.db.Create(&product)
	if result.Error != nil {
		return errors.New("failed to add product")
	}
	return nil
}

func (prd *ProductServiceImpl) GetAll() ([]*entity.Product, error) {
	//var product entity.Product
	var products []*entity.Product
	result := prd.db.Find(&products)
	if result.Error != nil {
		return nil, errors.New("couldn't find any product")
	}
	return products, nil
}

func (prd *ProductServiceImpl) GetById(id int) (*entity.Product, error) {
	var product entity.Product
	result := prd.db.First(&product, id)
	if result.Error != nil {
		return nil, errors.New("could not found")
	}
	return &product, nil

}

func (prd *ProductServiceImpl) Offer(productOfferReq *model.ProductOffer, userId int) error {

	product, err := prd.GetById(productOfferReq.ProductId)
	if err != nil {
		return err
	}
	if int(product.OfferPrice) > productOfferReq.OfferPrice {
		return errors.New("the offered price can't be lower than highest offer")
	}
	product.OfferPrice = uint(productOfferReq.OfferPrice)
	product.UserID = uint(userId)

	if result := prd.db.Save(&product); result.Error != nil {
		return result.Error
	}

	return nil
}
