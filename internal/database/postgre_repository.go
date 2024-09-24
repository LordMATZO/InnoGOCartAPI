package database

import (
	"database/sql"
	"fmt"
	"innogocartapi/internal/models"

	_ "github.com/lib/pq"
)

type postgreRepository struct {
	context *postgreContext
}

func (repository *postgreRepository) CartCreate(cart models.Cart) (models.Cart, error) {
	var resultCart models.Cart

	repository.context.db.QueryRowx("insert into cart (id)  Values(default)").Scan(resultCart)

	return resultCart, nil
}
func (repository *postgreRepository) ViewCart(cart models.Cart) (models.Cart, error) {
	var resultCartId models.CartId

	rows, error := (*repository).context.db.Query(fmt.Sprintf("select * from cart where id = %d", cart.Id))

	if error == nil {
		for rows.Next() {
			error = rows.Scan(&resultCartId)
			if error != nil {
				break
			}
		}
	}

	return models.Cart{Id: resultCartId}, error
}
func (repository *postgreRepository) DeleteFromCart(cartItem models.CartItem) (int64, error) {
	var sqlResult sql.Result
	var rowsAffected int64
	var error error

	sqlResult, error = repository.context.db.Exec(fmt.Sprintf("delete from cart_item where cart_id = %d and id = %d ", cartItem.Cart_id, cartItem.Id))

	if error == nil {
		rowsAffected, error = sqlResult.RowsAffected()
	}

	return rowsAffected, error
}
func (repository *postgreRepository) AddToCart(cartItem models.CartItem) (models.CartItem, error) {
	var resultCartItem models.CartItem
	db := repository.context.db

	rows := db.QueryRowx(fmt.Sprintf("insert into cart_item(cart_id, product, quantity) values(%d, '%s', %d)",
		cartItem.Cart_id,
		cartItem.Product,
		cartItem.Quantity))
	rows.Scan(&resultCartItem)

	return resultCartItem, nil
}

func newPostgreRepository(context DatabaseContext) *postgreRepository {
	repository := new(postgreRepository)

	repository.context = context.(*postgreContext)

	return repository
}
