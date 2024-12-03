package main

import (
	"desafio-ddd-go/infrastructure/database"
	"desafio-ddd-go/infrastructure/models"
	"desafio-ddd-go/usecase/order"
	prod "desafio-ddd-go/usecase/product"
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	ID    string  `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

type CreateOrderRequest struct {
	OrderID    string                   `json:"order_id" binding:"required"`
	CustomerID string                   `json:"customer_id" binding:"required"`
	Items      []CreateOrderItemRequest `json:"items" binding:"required,dive"`
}

type CreateOrderItemRequest struct {
	ID        string  `json:"id" binding:"required"`
	Name      string  `json:"name" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
}

func main() {

	database.Connect()

	server := gin.Default()

	server.POST("/create/product", func(ctx *gin.Context) {

		var req CreateProductRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Dados inválidos na requisição",
				"details": err.Error(),
			})
			return
		}

		product := prod.CreateProductDto{
			ID:    req.ID,
			Name:  req.Name,
			Price: req.Price,
		}

		usecase := prod.NewCreateProductUseCase(product, database.DB)

		err := usecase.Execute()
		if err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Erro ao executar usecase",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(201, gin.H{
			"message": "Produto criado com sucesso",
		})
	})

	server.POST("/create/order", func(ctx *gin.Context) {

		var req CreateOrderRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Dados inválidos na requisição",
				"details": err.Error(),
			})
			return
		}

		var customer models.CustomerModel
		if err := database.DB.First(&customer, "id = ?", req.CustomerID).Error; err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Cliente não encontrado",
				"details": err.Error(),
			})
			return
		}

		items := make([]order.CreateOrderItemDto, len(req.Items))
		for i, item := range req.Items {
			items[i] = order.CreateOrderItemDto{
				ID:        item.ID,
				Name:      item.Name,
				Price:     item.Price,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			}
		}

		dto := order.CreateOrderDto{
			OrderID:    req.OrderID,
			CustomerID: req.CustomerID,
			Items:      items,
		}

		usecase := order.NewCreateOrderUseCase(dto, database.DB)

		err := usecase.Execute()
		if err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Erro ao executar usecase",
				"details": err.Error(),
			})
			return
		}
		
		ctx.JSON(201, gin.H{
			"message": "Pedido criado com sucesso!",
		})

	})

	server.Run(":3000")
}

//o trecho comentado abaixo serve para criar um produto e um customer para testar essa rota de criação de order
/*
customer := &models.CustomerModel{
	ID:   "customer1",
	Name: "Test Customer",
}
if err := database.DB.Create(customer).Error; err != nil {
	log.Fatalf("Erro ao criar o cliente: %v", err)
}
*/
