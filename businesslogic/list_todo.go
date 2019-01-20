package businesslogic

import (
	"fmt"
	"github.com/pkprzekwas/gotodo/todos"
)

func ListTodo(todoService todos.TodoService) []todos.Todo {
	todoList, err := todoService.Todos()
	if err != nil {
		fmt.Println("Error listing todos.")
		fmt.Println(err)
	}
	return todoList
}
