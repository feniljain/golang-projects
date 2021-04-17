package cmd

import (
	"fmt"
	"todo/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			err := db.AddTask(arg)
			if err != nil {
				return
			}
		}
		fmt.Println("Task(s) added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
