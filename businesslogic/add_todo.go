package businesslogic

import "github.com/pkprzekwas/gotodo/todos"

func AddTodo(todoService todos.TodoService, title, description string) error {
	if err := todoService.CreateTodo(title, description); err != nil {
		return err
	}
	return nil
}
