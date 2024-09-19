package main

import (
	injector "innogocartapi/internal/app"
	"innogocartapi/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	injector.InitInjector()

	router := gin.Default()
	router.POST("/carts", services.CartCreate)
	router.POST("/carts/:cartid/items", services.AddToCart)
	router.DELETE("/carts/:cartid/items/:itemid", services.DeleteFromCart)
	router.GET("/carts/:cartid", services.ViewCart)

	router.Run("localhost:10101")
}
