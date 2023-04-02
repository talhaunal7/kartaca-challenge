package model

type ProductDto struct {
	ID            uint
	Name          string
	OfferPrice    uint
	UserID        *int
	ImgUrl        string
	UserFirstName string
	UserLastName  string
}
