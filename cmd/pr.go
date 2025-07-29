package cmd

import (
	"fmt"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
	"github.com/optiflowic/ghkit/internal/language"
	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/optiflowic/ghkit/internal/pr"
	"github.com/optiflowic/ghkit/internal/utils"
	"github.com/optiflowic/ghkit/internal/writer"
	"github.com/spf13/cobra"
)

type prCmdOptions struct {
	lang  string
	path  string
	force bool
}

func newPrCmd() *cobra.Command {
	opts := &prCmdOptions{}
	cmd := &cobra.Command{
		Use:   "pr",
		Short: "Add a pull request template to your repository",
		Long: `Add a GitHub pull request template to your repository.

This command downloads and installs a standardized pull request template.
You can specify the output path and the template language (such as 'en' or 'ja').

Examples:
  ghkit add pr --lang ja
  ghkit add pr --lang en --path ./repository
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewFromFlags(verbose, debug)
			f := fetcher.New(log)
			w := writer.New(log)
			c := commenter.New()
			service := pr.New(log, f, w, c)

			lang, err := language.New(opts.lang)
			if err != nil {
				return err
			}
			if !utils.Exists(opts.path) {
				return fmt.Errorf("the specified path does not exist: %s", opts.path)
			}

			err = service.Add(*lang, opts.path, opts.force)
			if err != nil {
				log.Error("Failed to add pr template", "error", err)
			}

			return nil
		},
	}

	cmd.Flags().
		StringVarP(&opts.lang, "lang", "l", "en", "The language of the template file (en or ja)")
	cmd.Flags().
		StringVar(&opts.path, "path", ".", "Output path for the template (default: current directory)")
	cmd.Flags().BoolVar(&opts.force, "force", false, "Overwrite file if it already exists")

	return cmd
}

func init() {
	addCmd.AddCommand(newPrCmd())
}
