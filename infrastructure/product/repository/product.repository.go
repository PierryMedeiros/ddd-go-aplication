package repository

import (
    "gorm.io/gorm"
    "desafio-ddd-go/domain/product/entity"
    orm "desafio-ddd-go/infrastructure/models"
)

type ProductRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(entity *entity.Product) error {
    productModel := orm.ProductModel{
        ID:    entity.GetID(),
        Name:  entity.GetName(),
        Price: entity.GetPrice(),
    }
    return r.db.Create(&productModel).Error
}

func (r *ProductRepository) Update(entity *entity.Product) error {
    productModel := orm.ProductModel{
        Name:  entity.GetName(),
        Price: entity.GetPrice(),
    }
    return r.db.Model(&orm.ProductModel{}).Where("id = ?", entity.GetID()).Updates(productModel).Error
}

func (r *ProductRepository) Find(id string) (*entity.Product, error) {
    var productModel orm.ProductModel
    if err := r.db.First(&productModel, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return entity.NewProduct(productModel.ID, productModel.Name, productModel.Price)
}

func (r *ProductRepository) FindAll() ([]*entity.Product, error) {
    var productModels []orm.ProductModel
    if err := r.db.Find(&productModels).Error; err != nil {
        return nil, err
    }

    products := make([]*entity.Product, len(productModels))
    for i, productModel := range productModels {
        product, err := entity.NewProduct(productModel.ID, productModel.Name, productModel.Price)
        if err != nil {
            return nil, err
        }
        products[i] = product
    }
    return products, nil
}
