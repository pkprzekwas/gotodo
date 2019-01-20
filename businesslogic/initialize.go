package businesslogic

import (
	"fmt"
	"github.com/pkprzekwas/gotodo/database"
	"github.com/pkprzekwas/gotodo/todos"
)

func Initialize(todoService *database.TodoService) {
	_ = database.CreateSchema(todoService.DB, todos.Schema)

	todoMap := map[string]string{
		"Dinner":      "Dinner in my parent's house.",
		"Bike repair": "Repairing wheel in the blue Trek bike.",
		"Cleaning":    "Cleaning up by bedroom before Rachel's visit.",
	}

	for title, description := range todoMap {
		err := todoService.CreateTodo(title, description)
		if err != nil {
			fmt.Println(title)
		}
	}
}
