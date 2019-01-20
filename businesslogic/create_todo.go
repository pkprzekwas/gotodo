package businesslogic

import (
	"fmt"
	"github.com/pkprzekwas/gotodo/todos"
)

func CreateTodo(todoService todos.TodoService, title, description string) {
	if err := todoService.CreateTodo(title, description); err != nil {
		fmt.Println("Error while creating ", title)
	}
}
