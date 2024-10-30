package entity

import "errors"

type ProductB struct {
    ID    string
    Name  string
    Price float64
}

func NewProductB(id string, name string, price float64) (*ProductB, error) {
    p := &ProductB{
        ID:    id,
        Name:  name,
        Price: price,
    }
    if err := p.Validate(); err != nil {
        return nil, err
    }
    return p, nil
}

func (p *ProductB) Validate() error {
    if len(p.ID) == 0 {
        return errors.New("id is required")
    }
    if len(p.Name) == 0 {
        return errors.New("name is required")
    }
    if p.Price < 0 {
        return errors.New("price must be greater than zero")
    }
    return nil
}

func (p *ProductB) ChangeName(name string) error {
    p.Name = name
    return p.Validate()
}

func (p *ProductB) ChangePrice(price float64) error {
    p.Price = price
    return p.Validate()
}

func (p *ProductB) GetID() string {
    return p.ID
}

func (p *ProductB) GetName() string {
    return p.Name
}

func (p *ProductB) GetPrice() float64 {
    return p.Price * 2
}
