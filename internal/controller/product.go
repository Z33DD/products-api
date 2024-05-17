package controller

import (
	"net/http"
	"strconv"

	"eulabs/internal/model"
	"eulabs/internal/service"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	db := service.GetDatabase()
	product := new(model.Product)
	err := c.Bind(product)

	if err != nil {
		// panic(err)
		return err
	}

	db.Create(product)
	return c.JSON(http.StatusOK, product)
}

func ReadProduct(c echo.Context) error {
	db := service.GetDatabase()
	id := c.Param("id")

	product := new(model.Product)
	db.First(&product, id)

	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	db := service.GetDatabase()

	id := c.Param("id")
	product := new(model.Product)

	err := c.Bind(product)
	if err != nil {
		return err
	}

	idStr, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	product.ID = uint(idStr)

	db.Save(&product)

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	db := service.GetDatabase()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	db.Delete(&model.Product{}, id)

	return c.JSON(http.StatusOK, "Product deleted")
}

func ReadAllProducts(c echo.Context) error {
	db := service.GetDatabase()
	products := []model.Product{}
	db.Find(&products)

	return c.JSON(http.StatusOK, products)
}
