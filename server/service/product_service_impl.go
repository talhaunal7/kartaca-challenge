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

	var product = entity.Product{
		Name:       productAddReq.Name,
		OfferPrice: productAddReq.OfferPrice,
		UserID:     nil,
	}
	result := prd.db.Create(&product)
	if result.Error != nil {
		return errors.New("failed to add product")
	}
	return nil
}

func (prd *ProductServiceImpl) GetAll() ([]*model.ProductDto, error) {
	var products []*entity.Product
	result := prd.db.Order("name desc").Preload("User").Find(&products)
	if result.Error != nil {
		return nil, errors.New("couldn't find any product")
	}

	var productDtos []*model.ProductDto
	for i, v := range products {
		productDto := model.ProductDto{
			ID:            v.ID,
			Name:          v.Name,
			OfferPrice:    v.OfferPrice,
			UserID:        v.UserID,
			ImgUrl:        v.ImgUrl,
			UserFirstName: v.User.FirstName,
			UserLastName:  v.User.LastName,
		}
		productDtos[i] = &productDto
	}

	return productDtos, nil
}

func (prd *ProductServiceImpl) Offer(productOfferReq *model.ProductOffer, userId int) error {

	product, err := prd.getById(productOfferReq.ProductId)
	if err != nil {
		return err
	}
	if int(product.OfferPrice) >= productOfferReq.OfferPrice {
		return errors.New("the offered price can't be equal or lower than highest offer")
	}
	product.OfferPrice = uint(productOfferReq.OfferPrice)
	product.UserID = &userId

	if result := prd.db.Save(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (prd *ProductServiceImpl) getById(id int) (*entity.Product, error) {
	var product entity.Product
	result := prd.db.First(&product, id)
	if result.Error != nil {
		return nil, errors.New("could not found")
	}
	return &product, nil
}
