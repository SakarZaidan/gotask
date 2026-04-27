package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Convert the ID from a String ("1") to an Integer (1).
		// Command line arguments always arrive as text.
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid task ID: %v\n", args[0])
			os.Exit(1)
		}

		// 2. Load all tasks
		tasks, err := store.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}

		// 3. Find the task with the matching ID
		found := false
		for i, t := range tasks {
			if t.ID == id {
				// Mark as complete and save
				tasks[i].Complete()
				found = true
				if err := store.Update(tasks[i]); err != nil {
					fmt.Fprintf(os.Stderr, "Error updating task: %v\n", err)
					os.Exit(1)
				}
				fmt.Printf("Marked task %d as done: %s\n", id, t.Title)
				break
			}
		}

		if !found {
			fmt.Fprintf(os.Stderr, "Task with ID %d not found\n", id)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
