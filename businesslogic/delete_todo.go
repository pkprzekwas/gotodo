package businesslogic

import "github.com/pkprzekwas/gotodo/todos"

func DeleteTodo(todoService todos.TodoService, id int) error {
	if err := todoService.DeleteTodo(id); err != nil {
		return err
	}
	return nil
}
