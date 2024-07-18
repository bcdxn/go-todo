package services

func NewStaticToDoService() ToDoService {
	return StaticToDoService{
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

type StaticToDoService struct {
	todos []ToDo
}

func (s StaticToDoService) GetToDo(id string) ToDo {
	return ToDo{
		Id:     "1",
		Task:   "Laundry",
		IsDone: false,
	}
}

func (s StaticToDoService) GetToDos() []ToDo {
	return s.todos
}
