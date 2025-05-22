package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Enter Project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enter Project v0.1 -- by DevShiba")
	},
}
