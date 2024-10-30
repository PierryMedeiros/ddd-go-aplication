package entity_test

import (
    "testing"
    "desafio-ddd-go/domain/checkout/entity"
    "github.com/stretchr/testify/assert"
)

func TestOrder_Validation(t *testing.T) {
    t.Run("should throw error when id is empty", func(t *testing.T) {
        items := []entity.OrderItem{}
        _, err := entity.NewOrder("", "123", items)
        assert.NotNil(t, err)
        assert.Equal(t, "id is required", err.Error())
    })

    t.Run("should throw error when customerId is empty", func(t *testing.T) {
        items := []entity.OrderItem{}
        _, err := entity.NewOrder("123", "", items)
        assert.NotNil(t, err)
        assert.Equal(t, "customerId is required", err.Error())
    })

    t.Run("should throw error when items are empty", func(t *testing.T) {
        items := []entity.OrderItem{}
        _, err := entity.NewOrder("123", "123", items)
        assert.NotNil(t, err)
        assert.Equal(t, "items are required", err.Error())
    })

    t.Run("should calculate total", func(t *testing.T) {
        item1 := entity.OrderItem{
            ID:        "i1",
            ProductID: "p1",
            Name:      "Item 1",
            Price:     100,
            Quantity:  2,
            Total:     200,
        }
        item2 := entity.OrderItem{
            ID:        "i2",
            ProductID: "p2",
            Name:      "Item 2",
            Price:     200,
            Quantity:  2,
            Total:     400,
        }
        order, err := entity.NewOrder("o1", "c1", []entity.OrderItem{item1})
        assert.Nil(t, err)
        assert.Equal(t, 200.0, order.Total)

        order2, err := entity.NewOrder("o1", "c1", []entity.OrderItem{item1, item2})
        assert.Nil(t, err)
        assert.Equal(t, 600.0, order2.Total)
    })

    t.Run("should throw error if the item quantity is less or equal to zero", func(t *testing.T) {
        item := entity.OrderItem{
            ID:        "i1",
            ProductID: "p1",
            Name:      "Item 1",
            Price:     100,
            Quantity:  0,
            Total:     0,
        }
        _, err := entity.NewOrder("o1", "c1", []entity.OrderItem{item})
        assert.NotNil(t, err)
        assert.Equal(t, "quantity must be greater than 0", err.Error())
    })
}
