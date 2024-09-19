package database

import "innogocartapi/internal/models"

type DatabaseRepository interface {
	CartCreate(models.Cart) (models.Cart, error)
	ViewCart(models.Cart) (models.Cart, error)
	DeleteFromCart(models.Cart) (models.Cart, error)
	AddToCart(models.Cart) (models.Cart, error)
}

func NewDatabaseRepository(context *databaseContext, databaseId int) DatabaseRepository {
	switch databaseId {
	case PostgreDatabaseId:
		return newPostgreRepository(*context)
	}

	return nil
}
