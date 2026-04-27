package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Fetch the list of tasks from storage
		tasks, err := store.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		// 2. Check if the user wants JSON output (Machine Readable)
		asJSON, _ := cmd.Flags().GetBool("json")
		if asJSON {
			data, _ := json.MarshalIndent(tasks, "", "  ")
			fmt.Println(string(data))
			return
		}

		// 3. User wants Pretty output (Human Readable)
		// We use tabwriter to create neatly aligned columns.
		// It automatically calculates the width needed for each column.
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		
		// Print the Header row
		fmt.Fprintln(w, "ID\tTITLE\tSTATUS\tCREATED")
		
		// Print each task row
		for _, t := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", 
				t.ID, 
				t.Title, 
				t.Status.String(), 
				t.CreatedAt.Format("2006-01-02 15:04"),
			)
		}
		
		// Flush() actually writes the accumulated data to the screen
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("json", false, "Output in JSON format")
}
