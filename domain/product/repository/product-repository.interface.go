package repository

import (
	"desafio-ddd-go/domain/1-shared/repository"
	entity "desafio-ddd-go/domain/product/entity"
)

type ProductRepositoryInterface interface {
	repository.RepositoryInterface[entity.Product]
}
