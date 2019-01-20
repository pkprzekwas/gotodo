package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/pkprzekwas/gotodo/todos"
)

var title, description string

func init() {
	createCmd.Flags().StringVar(
		&title,
		"title",
		"",
		"Title for a new Todo")
	createCmd.Flags().StringVar(
		&description,
		"description",
		"",
		"Description for a new Todo")
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create todo item.",
	Long:  `Create todo item by providing title and description.`,
	Run: func(cmd *cobra.Command, args []string) {
		var todoService todos.TodoService
		todoService = getTodoService()

		if err := businesslogic.CreateTodo(todoService, title, description); err != nil {
			fmt.Println(err)
		}
	},
}
