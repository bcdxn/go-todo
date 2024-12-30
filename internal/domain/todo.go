package domain

// ToDo represents a to-do item has been persisted.
type ToDo struct {
	ID     string
	IsDone bool
	Task   string
}

// ToDoNew represents a new to-do to be persisted.
type ToDoNew struct {
	Task string
}

// ToDoUpdate represents an update to a to-do being persisted.
type ToDoUpdate struct {
	Task   string
	IsDone bool
}
