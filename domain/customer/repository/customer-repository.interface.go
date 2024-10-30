package repository

import (
	"desafio-ddd-go/domain/1-shared/repository"
	entity "desafio-ddd-go/domain/customer/entity"
)

type CustomerRepositoryInterface interface {
	repository.RepositoryInterface[entity.Customer]
}
