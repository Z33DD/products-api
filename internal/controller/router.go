package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouterFactory() *echo.Echo {
	router := echo.New()

	router.POST("/product", CreateProduct)
	router.GET("/product/:id", ReadProduct)
	router.PUT("/product/:id", UpdateProduct)
	router.DELETE("/product/:id", DeleteProduct)
	router.GET("/product", ReadAllProducts)

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	return router
}
