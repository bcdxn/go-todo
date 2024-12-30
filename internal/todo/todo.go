// domain.go holds entities and interfaces that represent the ToDo domain

package todo

/* To-Do Domain Entities
------------------------------------------------------------------------------------------------- */

// ToDo represents a to-do item has been persisted.
type ToDo struct {
	ID     string
	Task   string
	IsDone bool
}
