package commands

import (
	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/pkprzekwas/gotodo/todos"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo items",
	Long:  `List todo items.`,
	Run: func(cmd *cobra.Command, args []string) {
		var todoService todos.TodoService
		todoService = getTodoService()
		_ = businesslogic.ListTodo(todoService)
	},
}
