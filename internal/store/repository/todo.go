package repository

import (
	"github.com/bcdxn/go-todo/internal/domain"
)

// ToDo is an interface that defines all functions that must be implemented to fulfill the
// obligations of the repository.
type ToDo interface {
	Create(domain.ToDoNew) (domain.ToDo, error)
	All() ([]domain.ToDo, error)
	GetByID(id string) (domain.ToDo, error)
	UpdateByID(domain.ToDoUpdate) (domain.ToDo, error)
	DeleteByID() (domain.ToDo, error)
}
