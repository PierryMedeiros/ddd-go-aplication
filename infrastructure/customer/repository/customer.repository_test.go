package repository_test

import (
	entity "desafio-ddd-go/domain/customer/entity"
	"desafio-ddd-go/domain/customer/value-object"
	"desafio-ddd-go/infrastructure/customer/repository"
	orm "desafio-ddd-go/infrastructure/models"
	"testing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	
	if err := db.AutoMigrate(&orm.CustomerModel{}); err != nil {
		return nil, err
	}
	return db, nil
}

func TestCustomerRepository_Create(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("could not set up database: %v", err)
	}
	repo := repository.NewCustomerRepository(db)

	addr := valueobject.Address{
		Street: "123 Main St",
		Number: 1,
		Zip:    "12345-678",
		City:   "Test City",
	}
	customer, err := entity.NewCustomer("1", "John Doe")
	if err != nil {
		t.Fatalf("could not create customer: %v", err)
	}
	customer.ChangeAddress(&addr)

	err = repo.Create(customer)
	if err != nil {
		t.Fatalf("could not create customer in db: %v", err)
	}
}

func TestCustomerRepository_Find(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("could not set up database: %v", err)
	}
	repo := repository.NewCustomerRepository(db)

	addr := valueobject.Address{
		Street: "123 Main St",
		Number: 1,
		Zip:    "12345-678",
		City:   "Test City",
	}
	customer, err := entity.NewCustomer("1", "John Doe")
	if err != nil {
		t.Fatalf("could not create customer: %v", err)
	}
	customer.ChangeAddress(&addr)
	repo.Create(customer)

	foundCustomer, err := repo.Find("1")
	if err != nil {
		t.Fatalf("could not find customer: %v", err)
	}

	if foundCustomer.GetId() != customer.GetId() {
		t.Errorf("expected id %s, got %s", customer.GetId(), foundCustomer.GetId())
	}
}

func TestCustomerRepository_Update(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("could not set up database: %v", err)
	}
	repo := repository.NewCustomerRepository(db)

	addr := valueobject.Address{
		Street: "123 Main St",
		Number: 1,
		Zip:    "12345-678",
		City:   "Test City",
	}
	customer, err := entity.NewCustomer("1", "John Doe")
	if err != nil {
		t.Fatalf("could not create customer: %v", err)
	}
	customer.ChangeAddress(&addr)
	repo.Create(customer)

	
	customer.ChangeName("Jane Doe")
	err = repo.Update(customer)
	if err != nil {
		t.Fatalf("could not update customer: %v", err)
	}

	
	updatedCustomer, err := repo.Find("1")
	if err != nil {
		t.Fatalf("could not find updated customer: %v", err)
	}

	if updatedCustomer.GetName() != "Jane Doe" {
		t.Errorf("expected name %s, got %s", "Jane Doe", updatedCustomer.GetName())
	}
}

func TestCustomerRepository_FindNotFound(t *testing.T) {
	db, err := setupDB()
	if err != nil {
		t.Fatalf("could not set up database: %v", err)
	}
	repo := repository.NewCustomerRepository(db)

	_, err = repo.Find("unknown")
	if err == nil || err.Error() != "customer not found" {
		t.Errorf("expected error 'customer not found', got %v", err)
	}
}
