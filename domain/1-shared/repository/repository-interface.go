package repository

type RepositoryInterface[T any] interface {
    Create(entity T) error
    update(id string) (*T, error)
    FindAll() ([]T, error)
    Find(id string) (*T, error)
}