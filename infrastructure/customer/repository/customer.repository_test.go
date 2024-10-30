package repository_test

import (
	"desafio-ddd-go/domain/customer/entity"
	"desafio-ddd-go/domain/customer/value-object"
    "gorm.io/driver/sqlite"
	"desafio-ddd-go/infrastructure/customer/repository"
	"desafio-ddd-go/infrastructure/customer/repository/orm"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setup() (*repository.CustomerRepository, func()) {
    dbConn, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        panic("Falha ao conectar com o banco de dados para os testes")
    }
    orm.Migrate(dbConn)
    repo := &repository.CustomerRepository{DB: dbConn}

    teardown := func() {
        dbConn.Exec("DELETE FROM customer_models")
    }

    return repo, teardown
}

func TestCustomerRepository_Create(t *testing.T) {
    repo, teardown := setup()
    defer teardown()

    customer, err := customer.NewCustomer("1", "John Doe")
    assert.NoError(t, err)

    address := valueobject.Address{
        Street:  "123 Main St",
        Number:  456,
        Zip: "12345",
        City:    "Cityville",
    }
    customer.ChangeAddress(&address)

    err = repo.Create(customer)
    assert.NoError(t, err)
}

func TestCustomerRepository_Find(t *testing.T) {
    repo, teardown := setup()
    defer teardown()

    customer, err := customer.NewCustomer("1", "John Doe")
    assert.NoError(t, err)

    address := valueobject.Address{
        Street:  "123 Main St",
        Number:  456,
        Zip: "12345",
        City:    "Cityville",
    }
    customer.ChangeAddress(&address)
    err = repo.Create(customer)
    assert.NoError(t, err)

    foundCustomer, err := repo.Find("1")
    assert.NoError(t, err)
    assert.Equal(t, "John Doe", foundCustomer.GetName())
    assert.Equal(t, "123 Main St", foundCustomer.GetAddress().Street)
}

func TestCustomerRepository_Update(t *testing.T) {
    repo, teardown := setup()
    defer teardown()

    customer, err := customer.NewCustomer("1", "John Doe")
    assert.NoError(t, err)

    address := valueobject.Address{
        Street:  "123 Main St",
        Number:  456,
        Zip: "12345",
        City:    "Cityville",
    }
    customer.ChangeAddress(&address)
    err = repo.Create(customer)
    assert.NoError(t, err)

    customer.ChangeName("Jane Doe")
    customer.AddRewardPoints(50)
    err = repo.Update(customer)
    assert.NoError(t, err)

    updatedCustomer, err := repo.Find("1")
    assert.NoError(t, err)
    assert.Equal(t, "Jane Doe", updatedCustomer.GetName())
    assert.Equal(t, 50, updatedCustomer.GetRewardPoints())
}

func TestCustomerRepository_FindAll(t *testing.T) {
    repo, teardown := setup()
    defer teardown()

    customer1, _ := customer.NewCustomer("1", "John Doe")
    address1 := valueobject.Address{
        Street:  "123 Main St",
        Number:  456,
        Zip: "12345",
        City:    "Cityville",
    }
    customer1.ChangeAddress(&address1)
    repo.Create(customer1)

    customer2, _ := customer.NewCustomer("2", "Jane Smith")
    address2 := valueobject.Address{
        Street:  "456 Oak St",
        Number:  789,
        Zip: "67890",
        City:    "Townsville",
    }
    customer2.ChangeAddress(&address2)
    repo.Create(customer2)

    customers, err := repo.FindAll()
    assert.NoError(t, err)
    assert.Equal(t, 2, len(customers))
}
