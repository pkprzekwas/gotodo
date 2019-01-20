package businesslogic

import (
	"fmt"
	"github.com/pkprzekwas/gotodo/todos"
)

func ListTodo(todoService todos.TodoService) []todos.Todo {
	todoList, err := todoService.Todos()
	if err != nil {
		fmt.Println("Error listing todo items.")
		fmt.Println(err)
	}

	for _, todo := range todoList {
		fmt.Println("___________________________", todo)
	}

	return todoList
}
