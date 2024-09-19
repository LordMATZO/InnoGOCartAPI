package models

type Product string
type CartItemId int

type CartItem struct {
	Id       CartItemId `json:"id"`
	Cart_id  CartId     `json:"cart_id"`
	Product  Product    `json:"product"`
	Quantity int        `json:"quantity"`
}
