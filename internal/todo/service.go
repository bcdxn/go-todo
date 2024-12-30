package todo

import "github.com/segmentio/ksuid"

type Service struct {
	Repository Repository
}

func (s Service) Create(in ServiceToDoCreateInput) (ToDo, error) {
	res, err := s.Repository.Create(RepositoryToDoCreateInput{
		ID:     ksuid.New().String(),
		Task:   in.Task,
		IsDone: false,
	})
	if err != nil {
		return ToDo{}, nil
	}

	return ToDo{
		ID:     res.ID,
		Task:   res.Task,
		IsDone: res.IsDone,
	}, nil
}

func (s Service) All() ([]ToDo, error) {
	var todos []ToDo
	rTodos, err := s.Repository.All()
	if err != nil {
		return todos, err
	}

	for _, t := range rTodos {
		todos = append(todos, ToDo{
			ID:     t.ID,
			Task:   t.Task,
			IsDone: t.IsDone,
		})
	}

	return todos, nil
}

func (s Service) GetByID(id string) (ToDo, error) {
	var t ToDo
	res, err := s.Repository.GetByID(id)
	if err != nil {
		return t, err
	}

	return ToDo{
		ID:     res.ID,
		Task:   res.Task,
		IsDone: res.IsDone,
	}, nil
}

func (s Service) UpdateByID(in ServiceToDoUpdateInput) (ToDo, error) {
	var t ToDo
	res, err := s.Repository.UpdateByID(RepositoryToDoUpdateInput{
		ID:     in.ID,
		Task:   in.Task,
		IsDone: in.IsDone,
	})
	if err != nil {
		return t, err
	}

	return ToDo{
		ID:     res.ID,
		Task:   res.Task,
		IsDone: res.IsDone,
	}, nil
}

func (s Service) DeleteByID(id string) (ToDo, error) {
	var t ToDo
	res, err := s.Repository.DeleteByID(id)
	if err != nil {
		return t, err
	}

	return ToDo{
		ID:     res.ID,
		Task:   res.Task,
		IsDone: res.IsDone,
	}, nil
}

type ServiceToDoCreateInput struct {
	Task string
}

type ServiceToDoUpdateInput struct {
	ID     string
	Task   string
	IsDone bool
}
