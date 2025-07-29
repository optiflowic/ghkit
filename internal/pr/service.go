package pr

import (
	"fmt"
	"path/filepath"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
	"github.com/optiflowic/ghkit/internal/format"
	"github.com/optiflowic/ghkit/internal/language"
	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/optiflowic/ghkit/internal/source_url"
	"github.com/optiflowic/ghkit/internal/utils"
	"github.com/optiflowic/ghkit/internal/writer"
)

type Service struct {
	log       logger.Logger
	fetcher   fetcher.Fetcher
	writer    writer.Writer
	commenter commenter.Commenter
}

func New(
	log logger.Logger,
	fetcher fetcher.Fetcher,
	writer writer.Writer,
	commenter commenter.Commenter,
) Service {
	return Service{
		log:       log,
		fetcher:   fetcher,
		writer:    writer,
		commenter: commenter,
	}
}

func (s Service) Add(
	lang language.Language,
	basePath string,
	force bool,
) error {
	filename := prTemplates[PullRequest]
	url := fmt.Sprintf(
		"%s/%s/.github/%s",
		source_url.Templates,
		lang,
		filename,
	)
	dest := filepath.Join(basePath, ".github", filename)
	if utils.Exists(dest) {
		if !force {
			s.log.Error("Already exists", "path", dest)
			return fmt.Errorf("already exists: %s", dest)
		}
		s.log.Warn("Already exists but will be overwritten by force", "path", dest)
	}

	data, err := s.fetcher.Fetch(url)
	if err != nil {
		s.log.Error("Failed to fetch", "url", url, "error", err)
		return fmt.Errorf("failed to fetch: %s", url)
	}

	content := s.commenter.PrependGeneratedComment(data, format.Markdown, url)
	if err := s.writer.Write(dest, content); err != nil {
		s.log.Error("Failed to write", "path", dest, "error", err)
		return fmt.Errorf("failed to write: %s", dest)
	}

	s.log.Info("Template added", "name", filename, "dest", dest)

	return nil
}
