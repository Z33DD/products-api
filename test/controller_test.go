package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"eulabs/internal/controller"
	"eulabs/internal/model"
	"eulabs/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up a test request
	product := model.Product{
		Name:  "Test Product",
		Price: 9.99,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the CreateProduct handler
	err := controller.CreateProduct(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response body
	var createdProduct model.Product
	err = json.Unmarshal(rec.Body.Bytes(), &createdProduct)
	assert.NoError(t, err)

	// Additional assertions on the created product
	assert.NotZero(t, createdProduct.ID)
	assert.Equal(t, product.Name, createdProduct.Name)
	assert.Equal(t, product.Price, createdProduct.Price)
}
func TestReadProduct(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up a test request
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()

	// Set up a test database
	db := service.GetDatabase()

	// Create a product for testing
	product := model.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       9.99,
	}
	db.Create(&product)

	// Call the ReadProduct handler
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(product.ID))) // Convert product.ID to string
	err := controller.ReadProduct(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response body
	var retrievedProduct model.Product
	err = json.Unmarshal(rec.Body.Bytes(), &retrievedProduct)
	assert.NoError(t, err)

	// Additional assertions on the retrieved product
	assert.Equal(t, product.ID, retrievedProduct.ID)
	assert.Equal(t, product.Name, retrievedProduct.Name)
	assert.Equal(t, product.Price, retrievedProduct.Price)
}
func TestDeleteProduct(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up a test request
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Set up a test database
	db := service.GetDatabase()

	// Create a product for testing
	product := model.Product{
		Name:  "Test Product",
		Price: 9.99,
	}
	db.Create(&product)

	// Call the DeleteProduct handler
	err := controller.DeleteProduct(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Check the response body
	expectedResponse := "\"Product deleted\"\n"
	assert.Equal(t, expectedResponse, rec.Body.String())
}
func TestUpdateProduct(t *testing.T) {
	// Set up a test database
	db := service.GetDatabase()

	// Create a product for testing
	existingProduct := model.Product{
		Name:  "Test Product",
		Price: 9.99,
	}
	db.Create(&existingProduct)

	// Create a new Echo instance
	e := echo.New()

	// Set up a test request
	product := model.Product{
		Name:  "Updated Product",
		Price: 19.99,
	}
	body, _ := json.Marshal(product)
	req := httptest.NewRequest(http.MethodPut, "/products/"+strconv.Itoa(int(existingProduct.ID)), bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(existingProduct.ID)))

	// Call the UpdateProduct handler
	err := controller.UpdateProduct(c)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response body
	var updatedProduct model.Product
	err = json.Unmarshal(rec.Body.Bytes(), &updatedProduct)
	assert.NoError(t, err)

	// Additional assertions on the updated product
	assert.Equal(t, existingProduct.ID, updatedProduct.ID, "%s != %s", existingProduct.ID, updatedProduct.ID)
	assert.Equal(t, product.Name, updatedProduct.Name)
	assert.Equal(t, product.Price, updatedProduct.Price)
}
