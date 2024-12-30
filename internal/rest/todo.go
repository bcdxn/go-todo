package rest

import "github.com/bcdxn/go-todo/internal/domain"

type toDo struct {
	ID     string `json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone`
}

func domainToDoFromRest(in toDo) domain.ToDo {
	return domain.ToDo{
		ID:     in.ID,
		Task:   in.Task,
		IsDone: in.IsDone,
	}
}

func restToDoFromDomain(in domain.ToDo) toDo {
	return toDo{
		ID:     in.ID,
		Task:   in.Task,
		IsDone: in.IsDone,
	}
}
