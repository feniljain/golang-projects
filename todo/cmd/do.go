package cmd

import (
	"fmt"
	"strconv"
	"todo/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse arguments")
				return
			}
			ids = append(ids, id)
		}
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Println("Something went wrong!")
			return
		}
		for _, id := range ids {
			if id < 0 || id > len(tasks) {
				fmt.Println("Invalid argument")
				continue
			} else {
				err := db.RemoveTask(id)
				if err != nil {
					fmt.Println("Error deleting task at", id)
				}
				fmt.Println("Successfully done task", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
