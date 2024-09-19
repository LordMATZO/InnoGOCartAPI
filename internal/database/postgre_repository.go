package database

import (
	"innogocartapi/internal/models"

	_ "github.com/lib/pq"
)

type postgreRepository struct {
	context postgreContext
}

func (repository postgreRepository) CartCreate(carts models.Cart) (models.Cart, error) {
	var resultCart models.Cart
	db := repository.context.db

	rows, _ := db.Query("select * from cart")

	for rows.Next() {
		error := rows.Scan(&resultCart)

		if error != nil {
			panic(error)
		}
	}

	return resultCart, nil
}
func (repository postgreRepository) ViewCart(carts models.Cart) (models.Cart, error) {
	var resultCart models.Cart
	db := repository.context.db

	rows, _ := db.Query("select * from cart")

	for rows.Next() {
		error := rows.Scan(&resultCart)

		if error != nil {
			panic(error)
		}
	}

	return resultCart, nil
}
func (repository postgreRepository) DeleteFromCart(carts models.Cart) (models.Cart, error) {
	var resultCart models.Cart
	db := repository.context.db

	rows, _ := db.Query("select * from cart")

	for rows.Next() {
		error := rows.Scan(&resultCart)

		if error != nil {
			panic(error)
		}
	}

	return resultCart, nil
}
func (repository postgreRepository) AddToCart(carts models.Cart) (models.Cart, error) {
	var resultCart models.Cart
	db := repository.context.db

	rows, _ := db.Query("select * from cart")

	for rows.Next() {
		error := rows.Scan(&resultCart)

		if error != nil {
			panic(error)
		}
	}

	return resultCart, nil
}

func newPostgreRepository(context databaseContext) postgreRepository {
	repository := postgreRepository{}

	repository.context = context.(postgreContext)

	return repository
}
