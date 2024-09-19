package services

import (
	"innogocartapi/internal/database"
	"innogocartapi/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func CartCreate(httpContext *gin.Context) {
	container := dig.New()
	var cartRepository database.DatabaseRepository

	error := container.Invoke(func(newCartRepository *database.DatabaseRepository) {
		cartRepository = *newCartRepository
	})

	if error != nil {
		panic(error)
	}

	resultCarts, error := cartRepository.CartCreate(models.Cart{})

	if error == nil {
		httpContext.IndentedJSON(http.StatusOK, resultCarts)
	} else {
		httpContext.IndentedJSON(http.StatusInternalServerError, resultCarts)
	}
}

func AddToCart(httpContext *gin.Context) {
	var resultCarts []models.CartItem
	//cartId := httpContext.Param("cartid")

	httpContext.IndentedJSON(http.StatusOK, resultCarts)
}

func DeleteFromCart(httpContext *gin.Context) {
	container := dig.New()
	var cartRepository database.DatabaseRepository
	//cartId := httpContext.Param("cartid")
	//itemId := httpContext.Param("itemid")

	error := container.Invoke(func(newCartRepository *database.DatabaseRepository) {
		cartRepository = *newCartRepository
	})

	if error != nil {
		panic(error)
	}

	resultCarts, error := cartRepository.DeleteFromCart(models.Cart{})

	if error == nil {
		httpContext.IndentedJSON(http.StatusOK, resultCarts)
	} else {
		httpContext.IndentedJSON(http.StatusInternalServerError, resultCarts)
	}
}

func ViewCart(httpContext *gin.Context) {
	container := dig.New()
	var cartRepository database.DatabaseRepository

	cartId := models.CartId(httpContext.Param("cartid"))

	error := container.Invoke(func(newCartRepository database.DatabaseRepository) {
		cartRepository = newCartRepository
	})

	if error != nil {
		panic(error)
	}

	resultCarts, error := cartRepository.ViewCart(models.Cart{Id: cartId})

	if error == nil {
		httpContext.IndentedJSON(http.StatusOK, resultCarts)
	} else {
		httpContext.IndentedJSON(http.StatusInternalServerError, resultCarts)
	}
}
