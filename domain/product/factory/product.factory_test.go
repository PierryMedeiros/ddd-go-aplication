package factory_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "desafio-ddd-go/domain/product/factory"
)

func TestProductFactory_Create(t *testing.T) {
    t.Run("Create Product A", func(t *testing.T) {
        product, err := factory.ProductFactory{}.Create("a", "Product A", 50.0)
        assert.Nil(t, err)
        assert.NotNil(t, product)
        assert.Equal(t, "Product A", product.GetName())
        assert.Equal(t, 50.0, product.GetPrice())
    })

    t.Run("Create Product B", func(t *testing.T) {
        product, err := factory.ProductFactory{}.Create("b", "Product B", 100.0)
        assert.Nil(t, err)
        assert.NotNil(t, product)
        assert.Equal(t, "Product B", product.GetName())
        assert.Equal(t, 200.0, product.GetPrice())
    })

    t.Run("Invalid Product Type", func(t *testing.T) {
        product, err := factory.ProductFactory{}.Create("c", "Product C", 150.0)
        assert.NotNil(t, err)
        assert.Nil(t, product)
        assert.Equal(t, "product type not supported", err.Error())
    })
}
