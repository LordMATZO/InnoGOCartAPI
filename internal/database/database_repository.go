package database

import (
	"innogocartapi/internal/models"
)

type DatabaseRepository interface {
	CartCreate(models.Cart) (models.Cart, error)
	ViewCart(models.Cart) (models.Cart, error)
	DeleteFromCart(models.CartItem) (int64, error)
	AddToCart(models.CartItem) (models.CartItem, error)
}

func NewDatabaseRepository(context DatabaseContext, databaseId int) DatabaseRepository {
	switch databaseId {
	case PostgreDatabaseId:
		return newPostgreRepository(context)
	}

	return nil
}
