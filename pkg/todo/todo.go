package todo

import "time"

type Service interface {
	GetToDo(id string) *ToDo
	GetToDos() []ToDo
}

type ToDo struct {
	Id        string
	Task      string
	IsDone    bool
	UpdatedAt time.Time
}
