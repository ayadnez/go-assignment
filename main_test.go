package main

import (
	"bytes"
	"encoding/json"
	"go-backend/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInsertProductHandler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()
	router.POST("/product", handler.InsertProduct)

	// Test case 1: Valid request
	validRequestBody := map[string]interface{}{
		"ProductName":             "Test Product",
		"ProductDescription":      "This is a test product",
		"ProductImage":            []string{"https://images.shiksha.ws/mediadata/images/articles/iStock-1144247639.jpg"},
		"ProductPrice":            19.99,
		"CompressedProductImages": []string{},
	}

	validRequestBodyBytes, _ := json.Marshal(validRequestBody)
	req := httptest.NewRequest("POST", "/product", bytes.NewBuffer(validRequestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "success", response["status"])
	assert.Equal(t, "inserted the product", response["result"])

	// Test case 2: Invalid request
	invalidRequestBody := map[string]interface{}{
		// Missing required fields
	}

	invalidRequestBodyBytes, _ := json.Marshal(invalidRequestBody)
	req = httptest.NewRequest("POST", "/product", bytes.NewBuffer(invalidRequestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Check the response status code for an invalid request
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Check the response body for an error response
	var errorResponse map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &errorResponse)

	assert.Equal(t, "Invalid request payload", errorResponse["error"])
}

func TestGetProductHandler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()
	router.GET("/product/:id", handler.GetProduct)

	// Test case 1: Product found
	req := httptest.NewRequest("GET", "/product/2", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "success", response["status"])
	assert.NotNil(t, response["result"])

	// Test case 2: Product not found
	req = httptest.NewRequest("GET", "/product/999", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Check the response status code for a not found scenario
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Check the response body for an error response
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "error", response["status"])
	assert.Equal(t, "product not found", response["message"])
}

func TestInsertUserHandler(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()
	router.POST("/user", handler.InsertUser)

	// Test case 1: Valid request
	validRequestBody := map[string]interface{}{
		"Id":        1,
		"Name":      "John Doe",
		"Mobile":    "1234567890",
		"Latitude":  40.7128,
		"Longitude": -74.0060,
	}

	validRequestBodyBytes, _ := json.Marshal(validRequestBody)
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(validRequestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "success", response["status:"])
	assert.Equal(t, "inserted the user", response["result:"])

}
