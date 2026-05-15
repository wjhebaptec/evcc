package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// These variables are set during build time via ldflags
var (
	Version = "0.0.1-dev"
	Commit  = "unknown"
	Date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version, commit hash, and build date of this evcc binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

// printVersion outputs the current version information to stdout
func printVersion() {
	fmt.Printf("evcc version %s\n", Version)
	fmt.Printf("  commit: %s\n", Commit)
	fmt.Printf("  built:  %s\n", Date)
}
