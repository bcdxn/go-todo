package model

import (
	"github.com/bcdxn/go-todo/internal/domain"
	"github.com/segmentio/ksuid"
)

var todosInMemory = []toDo{
	{
		ID:     "tod_" + ksuid.New().String(),
		Task:   "wash the dishes",
		IsDone: false,
	},
	{
		ID:     "tod_" + ksuid.New().String(),
		Task:   "fold laundry",
		IsDone: false,
	},
	{
		ID:     "tod_" + ksuid.New().String(),
		Task:   "create todo poc",
		IsDone: true,
	},
}

type ToDoInMemory struct{}

func (t ToDoInMemory) Create(domain.ToDoNew) (domain.ToDo, error) {
	return domain.ToDo{}, nil
}

func (t ToDoInMemory) All() ([]domain.ToDo, error) {
	var todos []domain.ToDo
	for _, item := range todosInMemory {
		todos = append(todos, domainToDoFromModel(item))
	}
	return todos, nil
}

func (t ToDoInMemory) GetByID(id string) (domain.ToDo, error) {
	return domain.ToDo{}, nil
}

func (t ToDoInMemory) UpdateByID(domain.ToDoUpdate) (domain.ToDo, error) {
	return domain.ToDo{}, nil
}

func (t ToDoInMemory) DeleteByID() (domain.ToDo, error) {
	return domain.ToDo{}, nil
}

func domainToDoFromModel(in toDo) domain.ToDo {
	return domain.ToDo{
		ID:     in.ID,
		Task:   in.Task,
		IsDone: in.IsDone,
	}
}

type toDo struct {
	ID     string
	Task   string
	IsDone bool
}
