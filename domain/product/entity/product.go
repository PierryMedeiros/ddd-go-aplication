package entity

import "errors"

type Product struct {
    ID    string
    Name  string
    Price float64
}

func NewProduct(id string, name string, price float64) (*Product, error) {
    p := &Product{
        ID:    id,
        Name:  name,
        Price: price,
    }
    if err := p.Validate(); err != nil {
        return nil, err
    }
    return p, nil
}

func (p *Product) Validate() error {
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

func (p *Product) ChangeName(name string) error {
    p.Name = name
    return p.Validate()
}

func (p *Product) ChangePrice(price float64) error {
    p.Price = price
    return p.Validate()
}

func (p *Product) GetID() string {
    return p.ID
}

func (p *Product) GetName() string {
    return p.Name
}

func (p *Product) GetPrice() float64 {
    return p.Price
}
