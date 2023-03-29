package model

type ProductOffer struct {
	ProductId  int `json:"productId" binding:"required"`
	OfferPrice int `json:"offerPrice" binding:"required"`
}
