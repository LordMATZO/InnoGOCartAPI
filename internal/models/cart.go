package models

type CartId int

type Cart struct {
	Id    CartId     `json:"id"`
	Items []CartItem `json:"items"`
}
