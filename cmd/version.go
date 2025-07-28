package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of ghkit",
	Long: `Print the current version of the ghkit CLI tool.

This version is set during build time. If you see 'dev', it means the binary was built without a version override.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
