package main

import (
	orderEntity "desafio-ddd-go/domain/checkout/entity"
	"desafio-ddd-go/domain/product/entity"
	"desafio-ddd-go/infrastructure/database"
	"desafio-ddd-go/infrastructure/models"
	orderRepo "desafio-ddd-go/infrastructure/order/repository"
	prodRepo "desafio-ddd-go/infrastructure/product/repository"
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

		repo := prodRepo.NewProductRepository(database.DB)
		product, err := entity.NewProduct(req.ID, req.Name, req.Price)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Erro ao criar o produto",
				"details": err.Error(),
			})
			return
		}

		err = repo.Create(product)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error":   "Erro ao salvar o produto no banco",
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

		var orderItems []orderEntity.OrderItem
		
		for _, item := range req.Items {
			var product models.ProductModel
			if err := database.DB.First(&product, "id = ?", item.ProductID).Error; err != nil {
				ctx.JSON(400, gin.H{
					"error":   "Produto não encontrado",
					"details": err.Error(),
				})
				return
			}

			orderItem := orderEntity.NewOrderItem(item.ID, item.Name, item.Price, item.ProductID, item.Quantity)
			orderItems = append(orderItems, *orderItem)
		}

		order, err := orderEntity.NewOrder(req.OrderID, req.CustomerID, orderItems)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error":   "Erro ao criar o pedido",
				"details": err.Error(),
			})
			return
		}

		repo := orderRepo.NewOrderRepository(database.DB)
		if err := repo.Create(order); err != nil {
			ctx.JSON(500, gin.H{
				"error":   "Erro ao salvar o pedido no banco",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(201, gin.H{
			"message": "Pedido criado com sucesso!",
		})

		//o trecho comentado abaixo serve para criar um produto e um customer para testar essa rota de criação de order
		/*customer := &models.CustomerModel{
			ID:   "customer1",
			Name: "Test Customer",
		}
		if err := database.DB.Create(customer).Error; err != nil {
			log.Fatalf("Erro ao criar o cliente: %v", err)
		}
		*/
	})

	server.Run(":3000")
}
