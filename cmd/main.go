package main

import (
	"desafio-ddd-go/domain/product/entity"
	"desafio-ddd-go/infrastructure/database"
	"desafio-ddd-go/infrastructure/models"
	"desafio-ddd-go/infrastructure/product/repository"
	"log"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	ID    string  `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func main() {

	database.Connect()

	if err := (&models.ProductModel{}).Migrate(database.DB); err != nil {
		log.Fatalf("Erro ao migrar: %v", err)
	}

	server := gin.Default()

	server.POST("/create/product", func(ctx *gin.Context) {

		repo := repository.NewProductRepository(database.DB)
		product, err := entity.NewProduct("22421", "Tefdgst Product", 100.0)

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

		var savedProduct models.ProductModel
		if err := database.DB.First(&savedProduct, "id = ?", product.ID).Error; err != nil {
			ctx.JSON(500, gin.H{
				"error":   "Erro ao verificar o produto no banco",
				"details": err.Error(),
			})
			return
		}

		ctx.JSON(201, gin.H{
			"message": "Produto criado com sucesso",
		})
	})

	server.GET("/product", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World222",
		})
	})

	server.Run(":3000")
}
