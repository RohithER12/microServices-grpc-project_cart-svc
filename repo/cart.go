package repo

import (
	"github.com/RohithER12/cart-svc/pkg/models"
	repoimpl "github.com/RohithER12/cart-svc/repo/repoImpl"
)

type Cart interface {
	CreateCart(cart models.Cart) error
	DeleteCart(id int64) error
	UpdateCart(cart models.Cart) error
	GetById(id int64) (models.Cart, error)
	GetByUserId(id int64) (models.Cart, error)
	DisplayCart(id int64) ([]models.CartItems, error)
}

func NewCartImpl() Cart {
	return &repoimpl.CartImpl{}
}
