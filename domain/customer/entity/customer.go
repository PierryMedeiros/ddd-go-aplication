package customer

import (
	"errors"
	valueobject "desafio-ddd-go/domain/customer/value-object"
)

type Customer struct {
    id            string
    name          string
    address       valueobject.Address
    active        bool
    rewardPoints  float64
}

func NewCustomer(id, name string) (*Customer, error) {
    c := &Customer{
        id:   id,
        name: name,
    }

    
    if err := c.Validate(); err != nil {
        return nil, err
    }

    return c, nil
}

func (c *Customer) Validate() error {
    if len(c.id) == 0 {
      return errors.New("id is required")
    }
    if len(c.name) == 0 {
      return errors.New("name is required")
    }
    return nil
}

func (c *Customer) ChangeName(name string) error {
    c.name = name
    return c.Validate()
}

func (c *Customer) ChangeAddress(address *valueobject.Address) {
    c.address = *address
}

func (c *Customer) Activate() error {
    if c.address == (valueobject.Address{}) {
        return errors.New("address is mandatory to activate a customer")
    }
    c.active = true
    return nil
}

func (c *Customer) Deactivate() {
    c.active = false
}

func (c *Customer) AddRewardPoints(points float64) {
    c.rewardPoints += points
}

func (c *Customer) GetId() string {
    return c.id
}

func (c *Customer) GetName() string {
    return c.name
}

func (c *Customer) GetAddress() valueobject.Address {
    return c.address
}

func (c *Customer) IsActive() bool {
    return c.active
}

func (c *Customer) GetRewardPoints() float64 {
    return c.rewardPoints
}
