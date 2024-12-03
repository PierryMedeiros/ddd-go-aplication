package product

import (
	"desafio-ddd-go/domain/product/entity"
	prodRepo "desafio-ddd-go/infrastructure/product/repository"
	"log"
	"gorm.io/gorm"
)

type CreateProductUseCase struct {
	ID string
	name string
	price float64
	db *gorm.DB
}

func NewCreateProductUseCase(props CreateProductDto, db *gorm.DB) *CreateProductUseCase {

	return &CreateProductUseCase{
		ID: props.ID,
		name: props.Name,
		price: props.Price,
		db: db,
	}
}

func (c *CreateProductUseCase) Execute() error{

	repo := prodRepo.NewProductRepository(c.db)
	product, err := entity.NewProduct(c.ID, c.name, c.price)
	if err != nil {
		log.Printf("Erro ao criar o produto: %v", err)
		return err
	}

	err = repo.Create(product)

	if err != nil {
		log.Printf("Erro ao criar inserir produto no banco: %v", err)
		return err
	}
	return nil
}
