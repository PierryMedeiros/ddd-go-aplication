package factory_test

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "desafio-ddd-go/domain/customer/factory"
    valueobject "desafio-ddd-go/domain/customer/value-object"
)

func TestCustomerFactory_Create(t *testing.T) {
    customer, err := factory.CustomerFactory{}.Create("John")
    assert.NoError(t, err)
    assert.NotNil(t, customer.GetId())
    assert.Equal(t, "John", customer.GetName())
    assert.Equal(t, valueobject.Address{}, customer.GetAddress())
}

func TestCustomerFactory_CreateWithAddress(t *testing.T) {
    address, err := valueobject.NewAddress("Street", 1, "13330-250", "São Paulo")
    assert.NoError(t, err)

    customer, err := factory.CustomerFactory{}.CreateWithAddress("John", *address)
    assert.NoError(t, err)
    assert.NotNil(t, customer.GetId())
    assert.Equal(t, "John", customer.GetName())
    assert.Equal(t, *address, customer.GetAddress())
}
