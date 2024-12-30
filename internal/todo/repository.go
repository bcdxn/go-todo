package todo

type Repository interface {
	Create(RepositoryToDoCreateInput) (RepositoryToDo, error)
	All() ([]RepositoryToDo, error)
	GetByID(id string) (RepositoryToDo, error)
	UpdateByID(RepositoryToDoUpdateInput) (RepositoryToDo, error)
	DeleteByID(id string) (RepositoryToDo, error)
}

type RepositoryToDoCreateInput struct {
	ID     string
	Task   string
	IsDone bool
}

type RepositoryToDoUpdateInput struct {
	ID     string
	Task   string
	IsDone bool
}

type RepositoryToDo struct {
	ID     string
	Task   string
	IsDone bool
}
