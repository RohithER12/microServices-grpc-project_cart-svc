package repo

import (
	"github.com/RohithER12/cart-svc/pkg/models"
	repoimpl "github.com/RohithER12/cart-svc/repo/repoImpl"
)

type CartItems interface {
	AddItem(item models.CartItems) error
	RemoveItem(item models.CartItems) error
	GetById(id int64) (models.CartItems, error)
	GetByCartId(id int64) ([]models.CartItems, error)
	// GetByUserId(id int64) ([]models.CartItems, error)
	RemoveOne(item models.CartItems) error
	GetByCartIdAndProductId(cartId, ProductId int64) (models.CartItems, error)
}

func NewCartItemsImpl() CartItems {
	return &repoimpl.CartItemImpl{}
}
