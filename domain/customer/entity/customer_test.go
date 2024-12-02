package customer_test

import (
    "testing"
    "desafio-ddd-go/domain/customer/entity"
    valueobject "desafio-ddd-go/domain/customer/value-object"
    "github.com/stretchr/testify/assert"
)

func TestShouldThrowErrorWhenIdIsEmpty(t *testing.T) {
    _, err := customer.NewCustomer("", "John")
    assert.EqualError(t, err, "id is required")
}

func TestShouldThrowErrorWhenNameIsEmpty(t *testing.T) {
    _, err := customer.NewCustomer("123", "")
    assert.EqualError(t, err, "name is required")
}

func TestShouldChangeName(t *testing.T) {
    c, err := customer.NewCustomer("123", "John")
    assert.NoError(t, err)
    err = c.ChangeName("Jane")
    assert.NoError(t, err)
    assert.Equal(t, "Jane", c.GetName())
}

func TestShouldActivate(t *testing.T) {
    c, err := customer.NewCustomer("1", "Customer 1")
    assert.NoError(t, err)

    address, err := valueobject.NewAddress("Street 1", 123, "13330-250", "Rio de Janeiro")
	if err != nil {
		t.Fatalf("Error creating address: %v", err)
	}
    c.ChangeAddress(address)

    err = c.Activate()
    assert.NoError(t, err)
    assert.True(t, c.IsActive())
}

func TestShouldThrowErrorWhenAddressIsUndefinedOnActivate(t *testing.T) {
    c, err := customer.NewCustomer("1", "Customer 1")
    assert.NoError(t, err)

    err = c.Activate()
    assert.EqualError(t, err, "address is mandatory to activate a customer")
}

func TestShouldDeactivate(t *testing.T) {
    c, err := customer.NewCustomer("1", "Customer 1")
    assert.NoError(t, err)

    c.Deactivate()
    assert.False(t, c.IsActive())
}

func TestShouldAddRewardPoints(t *testing.T) {
    c, err := customer.NewCustomer("1", "Customer 1")
    assert.NoError(t, err)

    assert.Equal(t, float64(0), c.GetRewardPoints())

    c.AddRewardPoints(10)
    assert.Equal(t, float64(10), c.GetRewardPoints())

    c.AddRewardPoints(10)
    assert.Equal(t, float64(20), c.GetRewardPoints())
}
