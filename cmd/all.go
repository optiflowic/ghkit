package cmd

import (
	"fmt"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
	"github.com/optiflowic/ghkit/internal/format"
	"github.com/optiflowic/ghkit/internal/issue"
	"github.com/optiflowic/ghkit/internal/language"
	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/optiflowic/ghkit/internal/meta"
	"github.com/optiflowic/ghkit/internal/pr"
	"github.com/optiflowic/ghkit/internal/utils"
	"github.com/optiflowic/ghkit/internal/writer"
	"github.com/spf13/cobra"
)

type allCmdOptions struct {
	format string
	lang   string
	path   string
	force  bool
}

func newAllCmd() *cobra.Command {
	opts := &allCmdOptions{}
	cmd := &cobra.Command{
		Use:   "all",
		Short: "Add all available GitHub templates",
		Long: `Add all GitHub default templates to your repository.

This command adds every template supported by ghkit, including issue templates,
pull request templates, and metadata files (e.g., CODEOWNERS, CONTRIBUTING).

Examples:
  ghkit add all
  ghkit add all --lang ja --path ./repository --format md
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewFromFlags(verbose, debug)

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

			f := fetcher.New(log)
			w := writer.New(log)
			c := commenter.New()
			issueService := issue.New(log, f, w, c)
			prService := pr.New(log, f, w, c)
			metaService := meta.New(log, f, w, c)

			err = issueService.Add(issue.All, *format, *lang, opts.path, opts.force)
			if err != nil {
				log.Error("Failed to add issue templates", "error", err)
			}
			err = prService.Add(*lang, opts.path, opts.force)
			if err != nil {
				log.Error("Failed to add pr template", "error", err)
			}
			err = metaService.Add(meta.All, *lang, opts.path, opts.force)
			if err != nil {
				log.Error("Failed to add meta templates", "error", err)
			}

			return nil
		},
	}

	cmd.Flags().
		StringVarP(&opts.format, "format", "f", "yml", "Format of the template file (yml or md)")
	cmd.Flags().
		StringVarP(&opts.lang, "lang", "l", "en", "The language of the template file (en or ja)")
	cmd.Flags().
		StringVar(&opts.path, "path", ".", "Output path for the template (default: current directory)")
	cmd.Flags().BoolVar(&opts.force, "force", false, "Overwrite file if it already exists")

	return cmd
}

func init() {
	addCmd.AddCommand(newAllCmd())
}
