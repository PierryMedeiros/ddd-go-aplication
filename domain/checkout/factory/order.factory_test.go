package factory_test

import (
    "testing"
    "desafio-ddd-go/domain/checkout/factory"
    "github.com/stretchr/testify/assert"
    "github.com/google/uuid"
)

func TestOrderFactory_Create(t *testing.T) {
    orderProps := factory.OrderFactoryProps{
        ID:         uuid.New().String(),
        CustomerID: uuid.New().String(),
        Items: []factory.OrderItemProps{
            {
                ID:        uuid.New().String(),
                Name:      "Product 1",
                ProductID: uuid.New().String(),
                Quantity:  1,
                Price:     100,
            },
        },
    }

    order, err := factory.OrderFactory{}.Create(orderProps)
    assert.Nil(t, err)
    assert.Equal(t, orderProps.ID, order.GetID())
    assert.Equal(t, orderProps.CustomerID, order.GetCustomerID())
    assert.Equal(t, len(orderProps.Items), len(order.GetItems()))
}
