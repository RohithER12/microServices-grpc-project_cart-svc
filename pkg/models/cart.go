package models

type Cart struct {
	Id          int64 `json:"id" gorm:"primaryKey"`
	UserId      int64 `json:"userId"`
	TotalAmount int64 `json:"totalAmnt"`
}

type CartItems struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	CartId    int64 `json:"cartId"`
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
	Amount    int64 `json:"amount"`
}
