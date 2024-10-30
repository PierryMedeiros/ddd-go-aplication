package factory

import (
    "github.com/google/uuid"
    "desafio-ddd-go/domain/customer/entity"
    valueobject "desafio-ddd-go/domain/customer/value-object"
)

type CustomerFactory struct{}

func (CustomerFactory) Create(name string) (*customer.Customer, error) {
    return customer.NewCustomer(uuid.New().String(), name)
}

func (CustomerFactory) CreateWithAddress(name string, address valueobject.Address) (*customer.Customer, error) {
    customer, err := customer.NewCustomer(uuid.New().String(), name)
    if err != nil {
        return nil, err
    }
    customer.ChangeAddress(&address)
    return customer, nil
}
