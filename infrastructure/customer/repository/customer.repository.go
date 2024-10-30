package repository

import (
	customerEntity "desafio-ddd-go/domain/customer/entity"
	valueobject "desafio-ddd-go/domain/customer/value-object"
	orm "desafio-ddd-go/infrastructure/models"
	"errors"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(entity *customerEntity.Customer) error {
	address := entity.GetAddress()
	customerModel := orm.CustomerModel{
		ID:           entity.GetId(),
		Name:         entity.GetName(),
		Street:       address.Street,
		Number:       address.Number,
		Zipcode:      address.Zip,
		City:         address.City,
		Active:       entity.IsActive(),
		RewardPoints: entity.GetRewardPoints(),
	}
	return r.db.Create(&customerModel).Error
}

func (r *CustomerRepository) Update(entity *customerEntity.Customer) error {
	address := entity.GetAddress()
	return r.db.Model(&orm.CustomerModel{}).Where("id = ?", entity.GetId()).Updates(orm.CustomerModel{
		Name:         entity.GetName(),
		Street:       address.Street,
		Number:       address.Number,
		Zipcode:      address.Zip,
		City:         address.City,
		Active:       entity.IsActive(),
		RewardPoints: entity.GetRewardPoints(),
	}).Error
}

func (r *CustomerRepository) Find(id string) (*customerEntity.Customer, error) {
	var customerModel orm.CustomerModel
	if err := r.db.First(&customerModel, "id = ?", id).Error; err != nil {
		return nil, errors.New("customer not found")
	}

	customer, err := customerEntity.NewCustomer(customerModel.ID, customerModel.Name)
	if err != nil {
		return nil, err
	}



    address := valueobject.Address{
        Street:  customerModel.Street,
        Number:  customerModel.Number,
        Zip: customerModel.Zipcode,
        City:    customerModel.City,
    }
    
    customer.ChangeAddress(&address)

	customer.ChangeAddress(&address)
	if customerModel.Active {
		customer.Activate()
	}
	customer.AddRewardPoints(customerModel.RewardPoints)

	return customer, nil
}

func (r *CustomerRepository) FindAll() ([]*customerEntity.Customer, error) {
	var customerModels []orm.CustomerModel
	if err := r.db.Find(&customerModels).Error; err != nil {
		return nil, err
	}

	var customers []*customerEntity.Customer
	for _, model := range customerModels {
		customer, err := customerEntity.NewCustomer(model.ID, model.Name)
		if err != nil {
			return nil, err
		}

		address := valueobject.Address{
            Street:  model.Street,
            Number:  model.Number,
            Zip: model.Zipcode,
            City:    model.City,
        }
		customer.ChangeAddress(&address)
		if model.Active {
			customer.Activate()
		}
		customer.AddRewardPoints(model.RewardPoints)

		customers = append(customers, customer)
	}

	return customers, nil
}
