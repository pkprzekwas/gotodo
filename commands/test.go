package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/pkprzekwas/gotodo/database"
	"github.com/pkprzekwas/gotodo/todos"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Simple integration test.",
	Long: `Temporary simple integration test
for the time of development.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Read config from config file and envs

		db, err := database.DBConnect(dbName, dbUser, dbPass)
		if err != nil {
			panic(err)
		}

		todoService := database.CreateTodoService(db)

		_ = database.CreateSchema(db, todos.Schema)

		todoMap := map[string]string{
			"Dinner":      "Dinner in my parent's house.",
			"Bike repair": "Repairing wheel in the blue Trek bike.",
			"Cleaning":    "Cleaning up by bedroom before Rachel's visit.",
		}

		for title, description := range todoMap {
			err = todoService.CreateTodo(title, description)
			if err != nil {
				fmt.Println(title)
			}
		}

		todo, err := todoService.Todos()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(todo)

		database.DropTable(db, "todo")
	},
}
