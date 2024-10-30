package factory

import (
    "github.com/google/uuid"
    "errors"
    "desafio-ddd-go/domain/product/entity"
)

type ProductFactory struct{}

func (ProductFactory) Create(typ string, name string, price float64) (entity.ProductInterface, error) {
    switch typ {
    case "a":
        return entity.NewProduct(uuid.New().String(), name, price)
    case "b":
        return entity.NewProductB(uuid.New().String(), name, price)
    default:
        return nil, errors.New("product type not supported")
    }
}
