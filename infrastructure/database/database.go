package database

import (
	"desafio-ddd-go/infrastructure/models"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPassword, dbName, dbPort)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Erro ao configurar conexão com o banco de dados: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	
	log.Println("Banco de dados conectado com sucesso!")

	if err := DB.AutoMigrate(
		&models.ProductModel{},
		&models.CustomerModel{},
		&models.OrderItemModel{},
		&models.OrderModel{},
	); err != nil {
		log.Fatalf("Erro ao migrar: %v", err)
	}
}
