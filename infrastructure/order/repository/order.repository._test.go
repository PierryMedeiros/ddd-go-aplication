package repository_test

import (
	"testing"

	"desafio-ddd-go/domain/checkout/entity"
	"desafio-ddd-go/infrastructure/models"
	"desafio-ddd-go/infrastructure/order/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&models.ProductModel{},
		&models.CustomerModel{},
		&models.OrderItemModel{},
		&models.OrderModel{},
	); err != nil {
		return nil, err
	}
	return db, nil
}

func TestOrderRepository_Create(t *testing.T) {
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("could not setup database: %v", err)
	}

	repo := repository.NewOrderRepository(db)
	orderItem := entity.NewOrderItem("item1", "Test Item", 10.0, "prod1", 2)
	order, err := entity.NewOrder("order1", "customer1", []entity.OrderItem{*orderItem})
	if err != nil {
		t.Fatalf("failed to create new order: %v", err)
	}

	t.Run("Create Order", func(t *testing.T) {
		err := repo.Create(order)
		assert.NoError(t, err)

		var savedOrder models.OrderModel
		err = db.Preload("Items").First(&savedOrder, "id = ?", order.GetID()).Error
		assert.NoError(t, err)
		assert.Equal(t, order.GetCustomerID(), savedOrder.CustomerID)
		assert.Equal(t, order.GetTotal(), savedOrder.Total)
		assert.Len(t, savedOrder.Items, 1)
	})

	t.Run("Update Order", func(t *testing.T) {
		orderItem2 := entity.NewOrderItem("item2", "Updated Item", 20.0, "prod2", 1)
		order.Items = append(order.Items, *orderItem2)

		err := repo.Update(order)
		assert.NoError(t, err)

		var updatedOrder models.OrderModel
		err = db.Preload("Items").First(&updatedOrder, "id = ?", order.GetID()).Error
		assert.NoError(t, err)
		assert.Equal(t, order.GetCustomerID(), updatedOrder.CustomerID)
		assert.Equal(t, order.GetTotal(), updatedOrder.Total)
		assert.Len(t, updatedOrder.Items, 2)
	})

	t.Run("Find Order", func(t *testing.T) {
		foundOrder, err := repo.Find(order.GetID())
		assert.NoError(t, err)
		assert.Equal(t, order.GetID(), foundOrder.GetID())
	})

	t.Run("Find All Orders", func(t *testing.T) {
		orders, err := repo.FindAll()
		assert.NoError(t, err)
		assert.Len(t, orders, 1)
	})
}
