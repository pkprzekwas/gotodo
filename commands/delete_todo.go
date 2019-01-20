package commands

import (
	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/pkprzekwas/gotodo/todos"
	"github.com/spf13/cobra"
)

var id int

func init() {
	deleteCmd.Flags().IntVar(
		&id,
		"id",
		-1,
		"Title for a new Todo")
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete todo item.",
	Long:  `Create todo item by providing title and description.`,
	Run: func(cmd *cobra.Command, args []string) {
		var todoService todos.TodoService
		todoService = getTodoService()
		businesslogic.DeleteTodo(todoService, id)
	},
}
