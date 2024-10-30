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
    rewardPoints  int
}

func NewCustomer(id, name string) (*Customer, error) {
    c := &Customer{
        id:   id,
        name: name,
    }

    // Valida o cliente no momento da criação
    if err := c.Validate(); err != nil {
        return nil, err
    }

    return c, nil
}

// Função para validar o cliente
func (c *Customer) Validate() error {
    if len(c.id) == 0 {
      return errors.New("id is required")
    }
    if len(c.name) == 0 {
      return errors.New("name is required")
    }
    return nil
}

// Método para alterar o nome do cliente
func (c *Customer) ChangeName(name string) error {
    c.name = name
    return c.Validate()
}

// Método para mudar o endereço do cliente
func (c *Customer) ChangeAddress(address *valueobject.Address) {
    c.address = *address
}

// Método para ativar o cliente
func (c *Customer) Activate() error {
    if c.address == (valueobject.Address{}) {
        return errors.New("address is mandatory to activate a customer")
    }
    c.active = true
    return nil
}

// Método para desativar o cliente
func (c *Customer) Deactivate() {
    c.active = false
}

// Método para adicionar pontos de recompensa ao cliente
func (c *Customer) AddRewardPoints(points int) {
    c.rewardPoints += points
}

// Getters
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

func (c *Customer) GetRewardPoints() int {
    return c.rewardPoints
}
