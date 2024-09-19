package models

type CartId string

type Cart struct {
	Id    CartId     `json:"id"`
	Items []CartItem `json:"items"`
}
