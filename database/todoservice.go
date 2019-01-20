package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkprzekwas/gotodo/todos"
)

type TodoService struct {
	db *sqlx.DB
}

func CreateTodoService(db *sqlx.DB) *TodoService {
	return &TodoService{db}
}

func (ts *TodoService) Todo(id int) (*todos.Todo, error) {
	var todoList []todos.Todo
	query := "SELECT * FROM todo WHERE id=$1"
	err := ts.db.Select(&todoList, query, id)
	if err != nil {
		return &todoList[0], err
	}
	return &todoList[0], nil
}

func (ts *TodoService) Todos() ([]todos.Todo, error) {
	var todoList []todos.Todo
	query := "SELECT * FROM todo"
	err := ts.db.Select(&todoList, query)
	if err != nil {
		return todoList, err
	}
	return todoList, nil
}

func (ts *TodoService) CreateTodo(title, description string) error {
	todo := &todos.Todo{
		Title:       title,
		Description: description,
		IsDone:      false,
	}

	tx := ts.db.MustBegin()

	query := "INSERT INTO todo (title, description, is_done) VALUES (:title, :description, :is_done)"
	if _, err := tx.NamedExec(query, todo); err != nil {
		return err
	}

	return tx.Commit()
}

func (ts *TodoService) DeleteTodo(id int) error {
	tx := ts.db.MustBegin()

	query := "DELETE FROM todo WHERE id=$1"
	_ = tx.MustExec(query, id)

	return tx.Commit()
}
