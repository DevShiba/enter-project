package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ep",
	Short: "ep is a CLI tool to quickly navigate to your project directories",
	Long: `Enter Project (ep) is a command-line interface tool
that helps you quickly change directories to your registered project folders.
You can add project parent paths and then use 'ep <project_name>' to jump there.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addPathCmd)

}
