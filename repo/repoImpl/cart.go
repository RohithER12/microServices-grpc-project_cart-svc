package repoimpl

import (
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
	var cart models.Cart
	if result := c.H.DB.Where("user_id = ?", userId).First(&cart); result.Error != nil {
		return models.Cart{}, result.Error
	}
	return cart, nil
}
