package cmd

import (
	"fmt"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
	"github.com/optiflowic/ghkit/internal/language"
	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/optiflowic/ghkit/internal/meta"
	"github.com/optiflowic/ghkit/internal/utils"
	"github.com/optiflowic/ghkit/internal/writer"
	"github.com/spf13/cobra"
)

type metaCmdOptions struct {
	lang  string
	path  string
	force bool
}

func newMetaCmd() *cobra.Command {
	opts := &metaCmdOptions{}
	cmd := &cobra.Command{
		Use:   "meta [name|all]",
		Short: "Add a metadata template (e.g., CODEOWNERS, CONTRIBUTING, or all)",
		Long: `Add GitHub metadata templates to your repository.

This command installs common repository meta files such as CODEOWNERS,
CONTRIBUTING.md, or all available templates. You can specify the output path,
language, and whether to overwrite existing files.

Examples:
  ghkit add meta all
  ghkit add meta codeowners --lang ja --path ./repository
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewFromFlags(verbose, debug)
			f := fetcher.New(log)
			w := writer.New(log)
			c := commenter.New()
			service := meta.New(log, f, w, c)

			template, err := meta.NewMetaTemplate(args[0])
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

			err = service.Add(*template, *lang, opts.path, opts.force)
			if err != nil {
				log.Error("Failed to add issue template", "error", err)
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
	addCmd.AddCommand(newMetaCmd())
}
