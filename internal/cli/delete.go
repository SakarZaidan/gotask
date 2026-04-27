package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Convert text ID to integer
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid task ID: %v\n", args[0])
			os.Exit(1)
		}

		// 2. Call the Delete method on our storage engine
		if err := store.Delete(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Deleted task: %d\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
