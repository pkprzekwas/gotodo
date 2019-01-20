package todos

import "fmt"

var Schema = `
CREATE TABLE todo (
    id SERIAL PRIMARY KEY,
    title TEXT,
	description TEXT,
    is_done BOOLEAN
)`

type Todo struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	IsDone      bool   `db:"is_done"`
}

func (t Todo) String() string {
	var status string
	if t.IsDone {
		status = "DONE"
	} else {
		status = "PENDING"
	}

	return fmt.Sprintf(
		"\nID: %d\nTitle: %s\nStatus: %s\nDescription:\n\t%s\n",
		t.ID,
		t.Title,
		status,
		t.Description)
}

type TodoService interface {
	Todo(id int) (*Todo, error)
	Todos() ([]Todo, error)
	CreateTodo(title, description string) error
	DeleteTodo(id int) error
}
