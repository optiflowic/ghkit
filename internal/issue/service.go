package issue

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
	f "github.com/optiflowic/ghkit/internal/format"
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
	template IssueTemplate,
	format f.Format,
	lang language.Language,
	basePath string,
	force bool,
) error {
	var filenames []string
	if template == All {
		files, err := all(format)
		if err != nil {
			s.log.Error("Could not find template files", "format", format, "error", err)
			return err
		}
		filenames = files
	} else {
		filename, err := find(template, format)
		if err != nil {
			s.log.Error("Could not find template file", "name", template, "format", format, "error", err)
			return err
		}
		filenames = []string{*filename}
	}

	var e error
	for _, filename := range filenames {
		url := fmt.Sprintf(
			"%s/%s/.github/ISSUE_TEMPLATE/%s",
			source_url.Templates,
			lang,
			filename,
		)
		dest := filepath.Join(basePath, ".github", "ISSUE_TEMPLATE", filename)
		if utils.Exists(dest) {
			if !force {
				s.log.Error("Already exists", "path", dest)
				e = errors.Join(e, fmt.Errorf("skipped because it already existed: %s", dest))
				continue
			}
			s.log.Warn("Already exists but will be overwritten by force", "path", dest)
		}

		data, err := s.fetcher.Fetch(url)
		if err != nil {
			s.log.Error("Failed to fetch", "url", url, "error", err)
			e = errors.Join(e, err)
			continue
		}

		content := s.commenter.PrependGeneratedComment(data, format, url)
		if err := s.writer.Write(dest, content); err != nil {
			s.log.Error("Failed to write", "path", dest, "error", err)
			e = errors.Join(e, err)
			continue
		}

		s.log.Info("Template added", "name", filename, "format", format, "dest", dest)
	}

	if e != nil {
		return e
	}

	return nil
}
