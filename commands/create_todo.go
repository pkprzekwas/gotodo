package commands

import (
	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/pkprzekwas/gotodo/todos"
	"github.com/spf13/cobra"
)

var title, description string

func init() {
	createCmd.Flags().StringVar(
		&title,
		"title",
		"",
		"Title for a new Todo")
	err := createCmd.MarkFlagRequired("title")
	if err != nil {
		panic(err)
	}
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
		businesslogic.CreateTodo(todoService, title, description)
	},
}
