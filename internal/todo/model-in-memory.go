package todo

import (
	"errors"

	"github.com/segmentio/ksuid"
)

var todos = []RepositoryToDo{
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

type ModelInMemory struct{}

func (m ModelInMemory) Create(t RepositoryToDoCreateInput) (RepositoryToDo, error) {
	newToDo := RepositoryToDo{
		ID:     t.ID,
		Task:   t.Task,
		IsDone: t.IsDone,
	}
	todos = append(todos, newToDo)
	return newToDo, nil
}

func (m ModelInMemory) GetByID(id string) (RepositoryToDo, error) {
	for _, t := range todos {
		if t.ID == id {
			return t, nil
		}
	}
	return RepositoryToDo{}, errors.New("NotFound")
}

func (m ModelInMemory) All() ([]RepositoryToDo, error) {
	return todos, nil
}

func (m ModelInMemory) UpdateByID(update RepositoryToDoUpdateInput) (RepositoryToDo, error) {
	for i, t := range todos {
		if t.ID == update.ID {
			todos[i].Task = update.Task
			todos[i].IsDone = update.IsDone
			return RepositoryToDo{
				ID:     update.ID,
				Task:   update.Task,
				IsDone: update.IsDone,
			}, nil
		}
	}
	return RepositoryToDo{}, errors.New("NotFound")
}

func (m ModelInMemory) DeleteByID(id string) (RepositoryToDo, error) {
	for i, t := range todos {
		if t.ID == id {
			// remove the element at i
			todos[i] = todos[len(todos)-1]
			todos = todos[:len(todos)-1]
			// return the deleted item
			return t, nil
		}
	}
	return RepositoryToDo{}, errors.New("NotFound")
}
