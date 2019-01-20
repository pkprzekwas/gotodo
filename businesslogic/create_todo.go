package businesslogic

import "github.com/pkprzekwas/gotodo/todos"

func CreateTodo(todoService todos.TodoService, title, description string) error {
	if err := todoService.CreateTodo(title, description); err != nil {
		return err
	}
	return nil
}
