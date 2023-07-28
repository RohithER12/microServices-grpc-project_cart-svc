package repoimpl

// import (
// 	"github.com/RohithER12/cart-svc/pkg/db"
// 	"github.com/RohithER12/cart-svc/pkg/models"
// )

// type CartItemImpl struct {
// 	H db.Handler
// }

// func (c *CartItemImpl) AddCartItem(item models.CartItems) error {
// 	if result := c.H.DB.Create(&item); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (c *CartItemImpl) RemoveCartItem(item models.CartItems) error {
// 	if result := c.H.DB.Delete(&item); result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func (c *CartItemImpl) GetCartItemById(id int64) (models.CartItems, error) {
// 	var item models.CartItems
// 	if result := c.H.DB.First(&item, id); result.Error != nil {
// 		return models.CartItems{}, result.Error
// 	}
// 	return item, nil
// }

// func (c *CartItemImpl) GetCartItemByCartId(cartId int64) ([]models.CartItems, error) {
// 	var items []models.CartItems
// 	if result := c.H.DB.Where("cart_id = ?", cartId).Find(&items); result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return items, nil
// }

// // func (c *CartItemImpl) GetByUserId(userId int64) ([]models.CartItems, error) {
// // 	var items []models.CartItems
// // 	if result := c.H.DB.Where("user_id = ?", userId).Find(&items); result.Error != nil {
// // 		return nil, result.Error
// // 	}
// // 	return items, nil
// // }

// // RemoveItem removes an item from the cart in the database.
// // If the quantity of the item is greater than 1, the quantity is decremented by 1.
// // If the quantity is exactly 1, the cart item is removed from the database.
// func (c *CartItemImpl) RemoveOneCartItem(item models.CartItems) error {
// 	var existingItem models.CartItems
// 	if result := c.H.DB.First(&existingItem, item.Id); result.Error != nil {
// 		return result.Error
// 	}

// 	if existingItem.Quantity > 1 {
// 		existingItem.Quantity--

// 		if result := c.H.DB.Save(&existingItem); result.Error != nil {
// 			return result.Error
// 		}
// 	} else {
// 		if result := c.H.DB.Delete(&item); result.Error != nil {
// 			return result.Error
// 		}
// 	}

// 	return nil
// }

// func (c *CartItemImpl) GetCartItemByCartIdAndProductId(cartId, productId int64) (models.CartItems, error) {
// 	var item models.CartItems
// 	if result := c.H.DB.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&item); result.Error != nil {
// 		return models.CartItems{}, result.Error
// 	}
// 	return item, nil
// }
