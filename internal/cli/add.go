package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/horux/gotask/internal/task"
	"github.com/spf13/cobra"
)

// addCmd defines the "add" subcommand.
// Example: "gotask add 'Buy eggs'"
var addCmd = &cobra.Command{
	Use:   "add [title]", // How the user should type it
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1), // Ensures the user provides at least the title
	Run: func(cmd *cobra.Command, args []string) {
		// args[0] is the title provided by the user
		title := args[0]
		
		// Retrieve the optional description from the --description flag
		description, _ := cmd.Flags().GetString("description")

		// Create a new Task object (In-memory)
		newTask := task.Task{
			Title:       title,
			Description: description,
			Status:      task.StatusPending,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Try to save the new task using our Store (JSON file)
		if err := store.Add(newTask); err != nil {
			// If there is an error, print it to Stderr (Standard Error) and exit
			fmt.Fprintf(os.Stderr, "Error adding task: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Added task: %s\n", title)
	},
}

func init() {
	// Register this command with the root "gotask" command
	rootCmd.AddCommand(addCmd)
	
	// Add a "Flag" specifically for this command. 
	// -d is the shorthand, --description is the full name.
	addCmd.Flags().StringP("description", "d", "", "Task description")
}
