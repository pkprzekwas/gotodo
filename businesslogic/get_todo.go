package businesslogic

import (
	"fmt"
	"github.com/pkprzekwas/gotodo/todos"
)

func GetTodo(todoService todos.TodoService, id int) *todos.Todo {
	todo, err := todoService.Todo(id)
	if err != nil {
		fmt.Println("Error getting todo: ", id)
		fmt.Println(err)
	}
	fmt.Println(todo)
	return todo
}
