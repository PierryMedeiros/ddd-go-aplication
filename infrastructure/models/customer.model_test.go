package models

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&CustomerModel{})
    return db, err
}

func TestCreateCustomer(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    customer := CustomerModel{
        ID:           "1",
        Name:         "John Doe",
        Street:       "123 Main St",
        Number:       123,
        Zipcode:      "12345",
        City:         "Anytown",
        Active:       true,
        RewardPoints: 100,
    }

    result := db.Create(&customer)
    assert.NoError(t, result.Error)
    assert.NotZero(t, customer.ID, "O ID do cliente deve ser diferente de zero após a criação")
}

func TestFindCustomer(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    customer := CustomerModel{
        ID:           "2",
        Name:         "Jane Doe",
        Street:       "456 Market St",
        Number:       456,
        Zipcode:      "67890",
        City:         "Othertown",
        Active:       false,
        RewardPoints: float64(50),
    }
    db.Create(&customer)

    var foundCustomer CustomerModel
    result := db.First(&foundCustomer, "id = ?", "2")

    assert.NoError(t, result.Error)
    assert.Equal(t, "Jane Doe", foundCustomer.Name)
    assert.Equal(t, "456 Market St", foundCustomer.Street)
    assert.Equal(t, 456, foundCustomer.Number)
    assert.Equal(t, "67890", foundCustomer.Zipcode)
    assert.Equal(t, "Othertown", foundCustomer.City)
    assert.False(t, foundCustomer.Active)
    assert.Equal(t, float64(50), foundCustomer.RewardPoints)
}

func TestUpdateCustomer(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    customer := CustomerModel{
        ID:           "3",
        Name:         "Mark Doe",
        Street:       "789 Elm St",
        Number:       789,
        Zipcode:      "11223",
        City:         "Oldtown",
        Active:       true,
        RewardPoints: 200,
    }
    db.Create(&customer)

    customer.Name = "Mark Updated"
    customer.City = "Newtown"
    result := db.Save(&customer)

    assert.NoError(t, result.Error)

    var updatedCustomer CustomerModel
    db.First(&updatedCustomer, "id = ?", "3")

    assert.Equal(t, "Mark Updated", updatedCustomer.Name)
    assert.Equal(t, "Newtown", updatedCustomer.City)
}

func TestDeleteCustomer(t *testing.T) {
    db, err := setupTestDB()
    assert.NoError(t, err)

    customer := CustomerModel{
        ID:           "4",
        Name:         "Eve Doe",
        Street:       "101 Pine St",
        Number:       101,
        Zipcode:      "33445",
        City:         "Tinyville",
        Active:       true,
        RewardPoints: float64(75),
    }
    db.Create(&customer)

    result := db.Delete(&customer)
    assert.NoError(t, result.Error)

    var deletedCustomer CustomerModel
    result = db.First(&deletedCustomer, "id = ?", "4")

    assert.Error(t, result.Error, "Cliente deletado não deve ser encontrado")
}
