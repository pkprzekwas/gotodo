package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/pkprzekwas/gotodo/todos"
)

var title, description string

func init() {
	addCmd.Flags().StringVar(
		&title,
		"title",
		"",
		"Title for a new Todo")
	addCmd.Flags().StringVar(
		&description,
		"description",
		"",
		"Description for a new Todo")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo item.",
	Long:  `Add todo item by providing title and description.`,
	Run: func(cmd *cobra.Command, args []string) {
		var todoService todos.TodoService
		todoService = getTodoService()

		if err := businesslogic.AddTodo(todoService, title, description); err != nil {
			fmt.Println(err)
		}
	},
}
