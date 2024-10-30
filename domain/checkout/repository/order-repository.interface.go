package repository

import (
    "desafio-ddd-go/domain/1-shared/repository"
	"desafio-ddd-go/domain/checkout/entity"
)

type OrderRepositoryInterface interface {
	repository.RepositoryInterface[entity.Order]
}
