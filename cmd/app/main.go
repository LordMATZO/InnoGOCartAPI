package main

import (
	"fmt"
	injector "innogocartapi/internal/app"
	"innogocartapi/internal/config"
	"innogocartapi/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration := config.NewCartConfiguration()

	injector.InitInjector()

	router := gin.Default()
	router.POST("/carts", services.CartCreate)
	router.POST("/carts/:cartid/items", services.AddToCart)
	router.DELETE("/carts/:cartid/items/:itemid", services.DeleteFromCart)
	router.GET("/carts/:cartid", services.ViewCart)

	router.Run(fmt.Sprintf("%s:%d", configuration.Server.Host, configuration.Server.Port))
}
