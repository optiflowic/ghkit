package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a GitHub project template to your repository",
	Long: `Add a predefined GitHub template file to your repository.

This command allows you to add common GitHub files such as issue templates, pull request templates, and metadata files.
Each subcommand specifies the type of template to add, such as 'issue', 'pr', or 'meta'.`,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
