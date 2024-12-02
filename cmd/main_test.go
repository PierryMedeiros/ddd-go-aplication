package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/create/product", func(ctx *gin.Context) {
		var req CreateProductRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Dados inválidos na requisição",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(201, gin.H{
			"message": "Produto criado com sucesso",
		})
	})

	r.POST("/create/order", func(ctx *gin.Context) {
		var req CreateOrderRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Dados inválidos na requisição",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(201, gin.H{
			"message": "Pedido criado com sucesso!",
		})
	})

	return r
}

func TestCreateProduct(t *testing.T) {
	r := setupRouter()

	product := CreateProductRequest{
		ID:    "prod1",
		Name:  "Produto 1",
		Price: 100.50,
	}

	productJSON, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/create/product", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "Produto criado com sucesso"}`, w.Body.String())
}

func TestCreateOrder(t *testing.T) {
	r := setupRouter()

	order := CreateOrderRequest{
		OrderID:    "order1",
		CustomerID: "customer1",
		Items: []CreateOrderItemRequest{
			{
				ID:        "item1",
				Name:      "Item 1",
				Price:     50.75,
				ProductID: "prod1",
				Quantity:  2,
			},
		},
	}

	orderJSON, _ := json.Marshal(order)

	req, _ := http.NewRequest("POST", "/create/order", bytes.NewBuffer(orderJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"message": "Pedido criado com sucesso!"}`, w.Body.String())
}
