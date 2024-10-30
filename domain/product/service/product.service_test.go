package service_test

import (
    "testing"
    "desafio-ddd-go/domain/product/entity"
    "desafio-ddd-go/domain/product/service"
    "github.com/stretchr/testify/assert"
)

func TestIncreasePrice(t *testing.T) {
    product1, _ := entity.NewProduct("1", "Product 1", 100.0)
    product2, _ := entity.NewProduct("2", "Product 2", 200.0)
    products := []*entity.Product{product1, product2}

    service.ProductService{}.IncreasePrice(products, 10)

    assert.Equal(t, 110.0, product1.GetPrice())
    assert.Equal(t, 220.0, product2.GetPrice())
}
