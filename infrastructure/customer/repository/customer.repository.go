package repository

import (
    "desafio-ddd-go/domain/customer/entity"
    "desafio-ddd-go/domain/customer/value-object"
    "desafio-ddd-go/infrastructure/customer/repository/orm"
    "errors"
    "gorm.io/gorm"
)

type CustomerRepository struct {
    DB *gorm.DB
}

func (r *CustomerRepository) Create(cust *customer.Customer) error {
    customerModel := orm.CustomerModel{
        ID:           cust.GetId(),
        Name:         cust.GetName(),
        Street:       cust.GetAddress().Street,
        Number:       cust.GetAddress().Number,
        Zipcode:      cust.GetAddress().Zip,
        City:         cust.GetAddress().City,
        Active:       cust.IsActive(),
        RewardPoints: cust.GetRewardPoints(),
    }
    return r.DB.Create(&customerModel).Error
}

func (r *CustomerRepository) Update(cust *customer.Customer) error {
    customerModel := orm.CustomerModel{
        ID:           cust.GetId(),
        Name:         cust.GetName(),
        Street:       cust.GetAddress().Street,
        Number:       cust.GetAddress().Number,
        Zipcode:      cust.GetAddress().Zip,
        City:         cust.GetAddress().City,
        Active:       cust.IsActive(),
        RewardPoints: cust.GetRewardPoints(),
    }
    return r.DB.Model(&orm.CustomerModel{}).Where("id = ?", cust.GetId()).Updates(&customerModel).Error
}

func (r *CustomerRepository) Find(id string) (*customer.Customer, error) {
    var customerModel orm.CustomerModel
    if err := r.DB.First(&customerModel, "id = ?", id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, errors.New("customer not found")
        }
        return nil, err
    }

    cust, err := customer.NewCustomer(customerModel.ID, customerModel.Name)
    if err != nil {
        return nil, err
    }

    address := valueobject.Address{
        Street:  customerModel.Street,
        Number:  customerModel.Number,
        Zip: customerModel.Zipcode,
        City:    customerModel.City,
    }
    cust.ChangeAddress(&address)

    if customerModel.Active {
        cust.Activate()
    }
    cust.AddRewardPoints(customerModel.RewardPoints)

    return cust, nil
}

func (r *CustomerRepository) FindAll() ([]*customer.Customer, error) {
    var customerModels []orm.CustomerModel
    if err := r.DB.Find(&customerModels).Error; err != nil {
        return nil, err
    }

    customers := make([]*customer.Customer, len(customerModels))
    for i, customerModel := range customerModels {
        cust, err := customer.NewCustomer(customerModel.ID, customerModel.Name)
        if err != nil {
            return nil, err
        }

        address := valueobject.Address{
            Street:  customerModel.Street,
            Number:  customerModel.Number,
            Zip: customerModel.Zipcode,
            City:    customerModel.City,
        }
        cust.ChangeAddress(&address)

        if customerModel.Active {
            cust.Activate()
        }
        cust.AddRewardPoints(customerModel.RewardPoints)

        customers[i] = cust
    }
    return customers, nil
}