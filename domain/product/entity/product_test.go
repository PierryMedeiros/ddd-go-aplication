package entity

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
    product, err := NewProduct("1", "Product 1", 100.0)
    assert.Nil(t, err)
    assert.Equal(t, "1", product.GetID())
    assert.Equal(t, "Product 1", product.GetName())
    assert.Equal(t, 100.0, product.GetPrice())
}

func TestChangeName(t *testing.T) {
    product, err := NewProduct("1", "Product 1", 100.0)
    assert.Nil(t, err)

    err = product.ChangeName("Product 2")
    assert.Nil(t, err)
    assert.Equal(t, "Product 2", product.GetName())
}

func TestChangePrice(t *testing.T) {
    product, err := NewProduct("1", "Product 1", 100.0)
    assert.Nil(t, err)

    err = product.ChangePrice(200.0)
    assert.Nil(t, err)
    assert.Equal(t, 200.0, product.GetPrice())
}

func TestProductValidation(t *testing.T) {
    _, err := NewProduct("", "Product 1", 100.0)
    assert.NotNil(t, err)
    assert.Equal(t, "id is required", err.Error())

    _, err = NewProduct("1", "", 100.0)
    assert.NotNil(t, err)
    assert.Equal(t, "name is required", err.Error())

    _, err = NewProduct("1", "Product 1", -10.0)
    assert.NotNil(t, err)
    assert.Equal(t, "price must be greater than zero", err.Error())
}
