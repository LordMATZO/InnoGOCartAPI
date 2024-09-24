package services

import (
	"fmt"
	injector "innogocartapi/internal/app"
	"innogocartapi/internal/database"
	"innogocartapi/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CartCreate(httpContext *gin.Context) {
	error := injector.DigContainer.Invoke(func(newCartRepository database.DatabaseRepository) {
		resultCarts, error := newCartRepository.CartCreate(models.Cart{})
		fmt.Println(error)
		if error == nil {
			httpContext.IndentedJSON(http.StatusOK, resultCarts)
		} else {
			httpContext.IndentedJSON(http.StatusInternalServerError, resultCarts)
		}
	})

	if error != nil {
		panic(error)
	}
}

func AddToCart(httpContext *gin.Context) {
	var error error
	var itemId int
	var cartId int
	var quantity int

	cartId, error = strconv.Atoi(httpContext.Param("cartid"))

	if error == nil {
		quantity, error = strconv.Atoi(httpContext.PostForm("quantity"))

		if error == nil {
			error = injector.DigContainer.Invoke(func(newCartRepository database.DatabaseRepository) {
				resultCartItem, error := newCartRepository.AddToCart(models.CartItem{
					Id:       models.CartItemId(itemId),
					Cart_id:  models.CartId(cartId),
					Product:  models.Product(httpContext.PostForm("product")),
					Quantity: quantity})

				if error == nil {
					httpContext.IndentedJSON(http.StatusOK, resultCartItem)
				} else {
					httpContext.IndentedJSON(http.StatusInternalServerError, resultCartItem)
				}
			})
		}
	}

	if error != nil {
		panic(error)
	}
}

func DeleteFromCart(httpContext *gin.Context) {
	var error error
	var itemId int
	var cartId int

	cartId, error = strconv.Atoi(httpContext.Param("cartid"))

	if error == nil {
		itemId, error = strconv.Atoi(httpContext.Param("itemid"))

		if error == nil {
			error = injector.DigContainer.Invoke(func(newCartRepository database.DatabaseRepository) {
				resultCartItem, error := newCartRepository.DeleteFromCart(models.CartItem{
					Id:      models.CartItemId(itemId),
					Cart_id: models.CartId(cartId)})

				if error == nil {
					httpContext.IndentedJSON(http.StatusOK, resultCartItem)
				} else {
					httpContext.IndentedJSON(http.StatusInternalServerError, resultCartItem)
				}
			})
		}
	}

	if error != nil {
		panic(error)
	}
}

func ViewCart(httpContext *gin.Context) {
	cartId, error := strconv.Atoi(httpContext.Param("cartid"))

	if error == nil {
		error = injector.DigContainer.Invoke(func(newCartRepository database.DatabaseRepository) {
			resultCarts, error := newCartRepository.ViewCart(models.Cart{Id: models.CartId(cartId)})
			fmt.Println(error)
			if error == nil {
				httpContext.IndentedJSON(http.StatusOK, resultCarts)
			} else {
				httpContext.IndentedJSON(http.StatusInternalServerError, error)
			}
		})
	}

	if error != nil {
		panic(error)
	}
}
