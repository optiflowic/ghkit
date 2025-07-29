package cmd

import (
	"fmt"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
	"github.com/optiflowic/ghkit/internal/format"
	"github.com/optiflowic/ghkit/internal/issue"
	"github.com/optiflowic/ghkit/internal/language"
	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/optiflowic/ghkit/internal/utils"
	"github.com/optiflowic/ghkit/internal/writer"
	"github.com/spf13/cobra"
)

type issueCmdOptions struct {
	format string
	lang   string
	path   string
	force  bool
}

func newIssueCmd() *cobra.Command {
	opts := &issueCmdOptions{}
	cmd := &cobra.Command{
		Use:   "issue [name|all]",
		Short: "Add an issue template (e.g., bug, feature, or all)",
		Long: `Add a GitHub issue template to your repository.

You can specify a single template (e.g., 'bug', 'feature') or use 'all' to add all available issue templates.
The output format can be either YAML (.yml) or Markdown (.md), and you can select a language such as English ('en') or Japanese ('ja').

Examples:
  ghkit add issue all --lang ja
  ghkit add issue bug --format md --lang en --path ./repository
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewFromFlags(verbose, debug)
			f := fetcher.New(log)
			w := writer.New(log)
			c := commenter.New()
			service := issue.New(log, f, w, c)

			template, err := issue.NewIssueTemplate(args[0])
			if err != nil {
				return err
			}
			format, err := format.New(opts.format)
			if err != nil {
				return err
			}
			lang, err := language.New(opts.lang)
			if err != nil {
				return err
			}
			if !utils.Exists(opts.path) {
				return fmt.Errorf("the specified path does not exist: %s", opts.path)
			}

			err = service.Add(*template, *format, *lang, opts.path, opts.force)
			if err != nil {
				log.Error("Failed to add issue template", "error", err)
			}

			return nil
		},
	}

	cmd.Flags().
		StringVarP(&opts.format, "format", "f", "yml", "Format of the template file (yml or md)")
	cmd.Flags().
		StringVarP(&opts.lang, "lang", "l", "en", "The language of the template file (en or ja)")
	cmd.Flags().
		StringVar(&opts.path, "path", ".", "Root directory path of the GitHub repository (default: current directory)")
	cmd.Flags().BoolVar(&opts.force, "force", false, "Overwrite file if it already exists")

	return cmd
}

func init() {
	addCmd.AddCommand(newIssueCmd())
}
