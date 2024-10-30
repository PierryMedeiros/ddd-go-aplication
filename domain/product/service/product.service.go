package service

import "desafio-ddd-go/domain/product/entity"

type ProductService struct{}

func (ProductService) IncreasePrice(products []*entity.Product, percentage float64) []*entity.Product {
    for _, product := range products {
        newPrice := (product.GetPrice() * percentage / 100) + product.GetPrice()
        product.ChangePrice(newPrice)
    }
    return products
}
