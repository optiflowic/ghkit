package cmd

import (
	"fmt"
	"strings"

	"github.com/optiflowic/ghkit/internal/issue"
	"github.com/optiflowic/ghkit/internal/meta"
	"github.com/optiflowic/ghkit/internal/pr"
	"github.com/spf13/cobra"
)

var templateType string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show available GitHub template files",
	Long: `List the available GitHub template files that can be added using the 'add' command.

You can filter by type using the --type (-t) flag:
  - issue: Issue templates (bug, feature, etc.)
  - pr: Pull request templates
  - meta: Meta files like CONTRIBUTING.md or CODEOWNERS
  - all: Show all templates (default)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		sections := []string{}
		switch templateType {
		case "all":
			sections = append(sections, formatSection("[📝 Issue Templates]", issue.ListAvailable()))
			sections = append(
				sections,
				formatSection("[📥 Pull Request Templates]", pr.ListAvailable()),
			)
			sections = append(sections, formatSection("[⚙️  Meta Files]", meta.ListAvailable()))
		case "issue":
			sections = append(sections, formatSection("[📝 Issue Templates]", issue.ListAvailable()))
		case "pr":
			sections = append(
				sections,
				formatSection("[📥 Pull Request Templates]", pr.ListAvailable()),
			)
		case "meta":
			sections = append(sections, formatSection("[⚙️  Meta Files]", meta.ListAvailable()))
		default:
			return fmt.Errorf("unsupported template type: %s", templateType)
		}

		fmt.Print("Available Templates:\n\n")
		fmt.Println(strings.Join(sections, "\n"))

		return nil
	},
}

func init() {
	listCmd.Flags().
		StringVarP(&templateType, "type", "t", "all", "The type of template you want to display (issue, pr, meta, or all)")
	rootCmd.AddCommand(listCmd)
}

func formatSection(title string, items []string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s\n", title))
	for _, item := range items {
		b.WriteString(fmt.Sprintf("  - %s\n", item))
	}
	return b.String()
}
