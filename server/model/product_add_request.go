package model

type ProductAdd struct {
	Name       string `json:"name" binding:"required"`
	OfferPrice uint   `json:"offerPrice" binding:"required"`
}
