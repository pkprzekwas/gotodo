package commands

import (
	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/pkprzekwas/gotodo/todos"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.Flags().IntVar(
		&id,
		"id",
		-1,
		"Todo id")
	err := getCmd.MarkFlagRequired("id")
	if err != nil {
		panic(err)
	}
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get todo item.",
	Long:  `Get todo item by providing id.`,
	Run: func(cmd *cobra.Command, args []string) {
		var todoService todos.TodoService
		todoService = getTodoService()
		businesslogic.GetTodo(todoService, id)
	},
}
