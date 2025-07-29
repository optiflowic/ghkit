package meta

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
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
	template MetaTemplate,
	lang language.Language,
	basePath string,
	force bool,
) error {
	var fileInfoList []FileInfo
	if template == All {
		fileInfoList = all()
	} else {
		fileInfo, err := find(template)
		if err != nil {
			s.log.Error("Could not find template file", "name", template, "error", err)
			return err
		}
		fileInfoList = []FileInfo{*fileInfo}
	}

	var e error
	for _, fileInfo := range fileInfoList {
		url := fmt.Sprintf(
			"%s/%s/.github/%s",
			source_url.Templates,
			lang,
			fileInfo.name,
		)
		dest := filepath.Join(basePath, ".github", fileInfo.name)
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

		content := s.commenter.PrependGeneratedComment(data, fileInfo.format, url)
		if err := s.writer.Write(dest, content); err != nil {
			s.log.Error("Failed to write", "path", dest, "error", err)
			e = errors.Join(e, err)
			continue
		}

		s.log.Info("Template added", "name", fileInfo.name, "dest", dest)
	}

	if e != nil {
		return e
	}

	return nil
}
