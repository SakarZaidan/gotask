/*
Package cli handles the User Interface (UI).
Since this is a CLI (Command Line Interface) app, the UI is text-based.
*/
package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/horux/gotask/internal/task"
	"github.com/spf13/cobra"
)

var (
	cfgFile string     // Path to the config/data file
	store   task.Store // The storage engine (using our Interface)
)

// rootCmd represents the base command when called without any subcommands.
// We use the "cobra" library which is the industry standard for Go CLI apps.
var rootCmd = &cobra.Command{
	Use:   "gotask",
	Short: "GoTask is a simple CLI task manager",
	Long: `A fully functional command-line task manager built with Go.
Users can create, list, update, complete, and delete tasks.
Data is persisted to a local JSON file.`,
}

// Execute is the main entry point for the CLI. It's called by main.go.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1) // Exit with 1 indicates an error happened
	}
}

// "init" is a special function in Go. It runs automatically BEFORE main().
func init() {
	// OnInitialize runs every time a command is executed.
	cobra.OnInitialize(initStore)

	// Add a "Flag" (like --config).
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotask.json)")
}

// initStore sets up our storage.
func initStore() {
	// 1. Check if user provided a file via the --config flag.
	if cfgFile == "" {
		// 2. If not, check if they set an Environment Variable (GOTASK_FILE).
		// Environment variables are a way to configure software without changing code.
		cfgFile = os.Getenv("GOTASK_FILE")
	}

	// 3. If still empty, use a default location in the user's Home directory.
	if cfgFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		cfgFile = filepath.Join(home, ".gotask.json")
	}

	// Initialize our JSON storage engine with the chosen file path.
	store = task.NewJSONStore(cfgFile)
}
