package main

import (
	"desafio-ddd-go/infrastructure/database"
	"desafio-ddd-go/infrastructure/models"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	if err := (&models.ProductModel{}).Migrate(database.DB); err != nil {
		log.Fatalf("Erro ao migrar: %v", err)
	}

	server := gin.Default()

	server.GET("/create/product", func(ctx *gin.Context) {
		product := models.ProductModel{
			ID:    "1", 
			Name:  "Produto Fictício",
			Price: 100.0,
		}

		if err := database.DB.Create(&product).Error; err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"message": "Produto criado com sucesso",
			"product": product,
		})
	})

	server.GET("/product", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Hello World222",
		})
	})

	server.Run(":3000")
}