package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"ep/internal/config"

	"github.com/spf13/cobra"
)

var addPathCmd = &cobra.Command{
	Use:   "add [path_to_add]",
	Short: "Add a new root path where your projects are located",
	Long: `Adds a new directory path to the list of paths Enter Project (ep)
will scan for projects. For example:
ep add /home/DevShiba/Development
ep add C:\Users\DevShiba\Projects`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pathToAd := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error loading configuration: %v\n", err)
			os.Exit(1)
		}

		absPathToAdd, err := filepath.Abs(pathToAd)
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error converting path to absolute: %v\n", err)
			os.Exit(1)
		}

		if slices.Contains(cfg.ProjectRoots, absPathToAdd) {
			fmt.Printf("Path '%s' already exists in configuration.\n", absPathToAdd)
			return
		}

		cfg.ProjectRoots = append(cfg.ProjectRoots, absPathToAdd)
		fmt.Printf("Added path: %s\n", absPathToAdd)

		if err := config.SaveConfig(cfg); err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "Error saving configuration: %v\n", err)
			os.Exit(1)
		}

	},
}
