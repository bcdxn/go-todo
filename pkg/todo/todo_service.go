package todo

import (
	"time"

	"github.com/hashicorp/go-hclog"
)

func NewStaticService(baseLogger hclog.Logger) Service {
	return &StaticService{
		todos: []ToDo{
			{
				Id:        "1",
				Task:      "Laundry",
				IsDone:    false,
				UpdatedAt: time.Now(),
			},
			{
				Id:        "2",
				Task:      "Dishes",
				IsDone:    false,
				UpdatedAt: time.Now(),
			},
		},
		l: baseLogger.Named("static_todo_service"),
	}
}

type StaticService struct {
	todos []ToDo
	l     hclog.Logger
}

func (s StaticService) GetToDo(id string) *ToDo {
	s.l.Trace("getting todo by ID", "fn", "GetToDo")

	for _, t := range s.todos {
		if t.Id == id {
			return &t
		}
	}

	return nil
}

func (s *StaticService) GetToDos() []ToDo {
	s.l.Trace("getting all todos", "fn", "GetToDos")
	return s.todos
}
