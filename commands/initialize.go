package commands

import (
	"github.com/pkprzekwas/gotodo/businesslogic"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize database.",
	Long:  `Initialize database and populate with test data.`,
	Run: func(cmd *cobra.Command, args []string) {
		businesslogic.Initialize(getTodoService())
	},
}
