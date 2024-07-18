package services

type ToDoService interface {
	GetToDo(id string) ToDo
	GetToDos() []ToDo
}

type ToDo struct {
	Id     string
	Task   string
	IsDone bool
}
