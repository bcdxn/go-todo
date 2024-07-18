package services

import (
	"errors"
)

type ToDoService struct {
	todos []ToDo
}

func NewToDoService() *ToDoService {
	return &ToDoService{
		todos: []ToDo{
			{
				Id:     "1",
				Task:   "Laundry",
				IsDone: false,
			},
			{
				Id:     "2",
				Task:   "Dishes",
				IsDone: false,
			},
		},
	}
}

func (s ToDoService) GetToDos() []ToDo {
	return s.todos
}

func (s *ToDoService) AddToDo(t ToDo) {
	s.todos = append(s.todos, t)
}

func (s *ToDoService) Update(updatedTodo ToDo) error {
	for i, todo := range s.todos {
		if todo.Id == updatedTodo.Id {
			s.todos[i] = updatedTodo
			return nil
		}
	}

	return errors.New("TODO not found")
}

type ToDo struct {
	Id     string
	Task   string
	IsDone bool
}

type GetToDoService interface {
	GetToDos() []ToDo
}
