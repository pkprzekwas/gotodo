package commands

import (
	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/pkprzekwas/gotodo/todos"
	"github.com/spf13/cobra"
)

func init() {
	deleteCmd.Flags().IntVar(
		&id,
		"id",
		-1,
		"Todo id")
	err := deleteCmd.MarkFlagRequired("id")
	if err != nil {
		panic(err)
	}
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete todo item.",
	Long:  `Delete todo item by providing id.`,
	Run: func(cmd *cobra.Command, args []string) {
		var todoService todos.TodoService
		todoService = getTodoService()
		businesslogic.DeleteTodo(todoService, id)
	},
}
