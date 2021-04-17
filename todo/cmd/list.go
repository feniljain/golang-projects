package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Something went wrong")
		}
		for _, task := range tasks {
			fmt.Println(task)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
