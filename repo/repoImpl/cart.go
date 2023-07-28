package repoimpl

import (
	"fmt"

	"github.com/RohithER12/cart-svc/pkg/db"
	"github.com/RohithER12/cart-svc/pkg/models"
)

type CartImpl struct {
	H db.Handler
}

func (c *CartImpl) CreateCart(cart models.Cart) error {

	if result := c.H.DB.Create(&cart); result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CartImpl) GetById(id int64) (models.Cart, error) {
	var cart models.Cart
	if result := c.H.DB.First(&cart, id); result.Error != nil {
		return models.Cart{}, result.Error
	}

	return cart, nil
}

func (c *CartImpl) DisplayCart(id int64) ([]models.CartItems, error) {
	cart := []models.CartItems{}
	if result := c.H.DB.Find(&cart); result.Error != nil {
		return nil, result.Error
	}
	return cart, nil

}

func (c *CartImpl) DeleteCart(id int64) error {
	if result := c.H.DB.Delete(&models.Cart{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *CartImpl) UpdateCart(cart models.Cart) error {
	if result := c.H.DB.Save(&cart); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *CartImpl) GetByUserId(userId int64) (models.Cart, error) {
	fmt.Println("reppppppppppppppppppooooooooooooooooo", userId)
	var cart models.Cart
	if result := c.H.DB.Where("user_id = ?", userId).First(&cart); result.Error != nil {
		return models.Cart{}, result.Error
	}
	return cart, nil
}
func (c *CartImpl) AddCartItem(item models.CartItems) error {
	if result := c.H.DB.Create(&item); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *CartImpl) RemoveCartItem(item models.CartItems) error {
	if result := c.H.DB.Delete(&item); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c *CartImpl) GetCartItemById(id int64) (models.CartItems, error) {
	var item models.CartItems
	if result := c.H.DB.First(&item, id); result.Error != nil {
		return models.CartItems{}, result.Error
	}
	return item, nil
}

func (c *CartImpl) GetCartItemByCartId(cartId int64) ([]models.CartItems, error) {
	var items []models.CartItems
	if result := c.H.DB.Where("cart_id = ?", cartId).Find(&items); result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

// func (c *CartItemImpl) GetByUserId(userId int64) ([]models.CartItems, error) {
// 	var items []models.CartItems
// 	if result := c.H.DB.Where("user_id = ?", userId).Find(&items); result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return items, nil
// }

// RemoveItem removes an item from the cart in the database.
// If the quantity of the item is greater than 1, the quantity is decremented by 1.
// If the quantity is exactly 1, the cart item is removed from the database.
func (c *CartImpl) RemoveOneCartItem(item models.CartItems) error {
	var existingItem models.CartItems
	if result := c.H.DB.First(&existingItem, item.Id); result.Error != nil {
		return result.Error
	}

	if existingItem.Quantity > 1 {
		existingItem.Quantity--

		if result := c.H.DB.Save(&existingItem); result.Error != nil {
			return result.Error
		}
	} else {
		if result := c.H.DB.Delete(&item); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (c *CartImpl) GetCartItemByCartIdAndProductId(cartId, productId int64) (models.CartItems, error) {
	var item models.CartItems
	if result := c.H.DB.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&item); result.Error != nil {
		return models.CartItems{}, result.Error
	}
	return item, nil
}
