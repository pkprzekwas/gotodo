package businesslogic

import (
	"fmt"
	"github.com/pkprzekwas/gotodo/todos"
)

func DeleteTodo(todoService todos.TodoService, id int) {
	if err := todoService.DeleteTodo(id); err != nil {
		fmt.Println("Error while deleting ", id)
	}
}
