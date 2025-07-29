package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool
	debug   bool
)

var rootCmd = &cobra.Command{
	Use:   "ghkit",
	Short: "CLI tool for adding standard GitHub project templates",
	Long: `ghkit is a command-line tool that helps you add common GitHub project templates such as:

  - Issue templates (YAML/Markdown)
  - Pull request templates
  - Contribution guidelines
  - CODEOWNERS, SECURITY.md, SUPPORT.md, and other meta files

Templates can be filtered and added selectively or all at once.
You can also customize language, output path, and verbosity.

Typical usage:

  ghkit add all
  ghkit add issue bug --format md --lang ja
  ghkit list
  ghkit version
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug logging")
}
