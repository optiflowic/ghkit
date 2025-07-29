package meta

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/optiflowic/ghkit/internal/commenter"
	"github.com/optiflowic/ghkit/internal/fetcher"
	"github.com/optiflowic/ghkit/internal/format"
	"github.com/optiflowic/ghkit/internal/language"
	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/optiflowic/ghkit/internal/source_url"
	"github.com/optiflowic/ghkit/internal/writer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Add(t *testing.T) {
	tmp := t.TempDir()
	log := logger.NewWithWriter(io.Discard, logger.LevelError)

	t.Run("single template success", func(t *testing.T) {
		lang := language.English
		filename := "CODEOWNERS"
		url := fmt.Sprintf(
			"%s/%s/.github/%s",
			source_url.Templates,
			lang,
			filename,
		)
		data := []byte("template data")
		exceptedPath := filepath.Join(tmp, ".github", filename)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		f := fetcher.NewMockFetcher(ctrl)
		w := writer.NewMockWriter(ctrl)
		c := commenter.New()

		f.EXPECT().Fetch(url).Return(data, nil).Times(1)
		w.EXPECT().
			Write(exceptedPath, c.PrependGeneratedComment(data, format.PlaneText, url)).
			Return(nil).
			Times(1)
		service := New(log, f, w, c)

		err := service.Add(CodeOwners, lang, tmp, false)

		assert.NoError(t, err)
	})

	t.Run("all templates success", func(t *testing.T) {
		lang := language.English
		fileInfoList := all()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		f := fetcher.NewMockFetcher(ctrl)
		w := writer.NewMockWriter(ctrl)
		c := commenter.New()

		for _, fileInfo := range fileInfoList {
			url := fmt.Sprintf(
				"%s/%s/.github/%s",
				source_url.Templates,
				lang,
				fileInfo.name,
			)
			data := []byte(fmt.Sprintf("%s template data", fileInfo.name))
			exceptedPath := filepath.Join(tmp, ".github", fileInfo.name)
			f.EXPECT().Fetch(url).Return(data, nil).Times(1)
			w.EXPECT().
				Write(exceptedPath, c.PrependGeneratedComment(data, fileInfo.format, url)).
				Return(nil).
				Times(1)
		}
		service := New(log, f, w, c)

		err := service.Add(All, lang, tmp, false)

		assert.NoError(t, err)
	})

	t.Run("already exists skipped", func(t *testing.T) {
		filename := "CODEOWNERS"
		existingPath := filepath.Join(tmp, ".github", filename)
		assert.NoError(t, os.MkdirAll(filepath.Dir(existingPath), 0755))
		assert.NoError(t, os.WriteFile(existingPath, []byte("existing"), 0644))
		defer func() {
			err := os.RemoveAll(existingPath)
			assert.NoError(t, err)
		}()

		lang := language.English
		url := fmt.Sprintf(
			"%s/%s/.github/%s",
			source_url.Templates,
			lang,
			filename,
		)
		data := []byte("template data")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		f := fetcher.NewMockFetcher(ctrl)
		w := writer.NewMockWriter(ctrl)
		c := commenter.New()

		f.EXPECT().Fetch(url).Return(data, nil).Times(0)
		w.EXPECT().
			Write(existingPath, c.PrependGeneratedComment(data, format.PlaneText, url)).
			Return(nil).
			Times(0)
		service := New(log, f, w, c)

		err := service.Add(CodeOwners, lang, tmp, false)

		assert.Error(t, err)
	})

	t.Run("already exists force add", func(t *testing.T) {
		filename := "CONTRIBUTING.md"
		existingPath := filepath.Join(tmp, ".github", filename)
		assert.NoError(t, os.MkdirAll(filepath.Dir(existingPath), 0755))
		assert.NoError(t, os.WriteFile(existingPath, []byte("existing"), 0644))
		defer func() {
			err := os.RemoveAll(existingPath)
			assert.NoError(t, err)
		}()

		lang := language.English
		url := fmt.Sprintf(
			"%s/%s/.github/%s",
			source_url.Templates,
			lang,
			filename,
		)
		data := []byte("template data")

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		f := fetcher.NewMockFetcher(ctrl)
		w := writer.NewMockWriter(ctrl)
		c := commenter.New()

		content := c.PrependGeneratedComment(data, format.Markdown, url)
		f.EXPECT().Fetch(url).Return(data, nil).Times(1)
		w.EXPECT().Write(existingPath, content).Return(nil).Times(1)
		service := New(log, f, w, c)

		err := service.Add(Contributing, lang, tmp, true)

		assert.NoError(t, err)
	})

	t.Run("invalid template", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		f := fetcher.NewMockFetcher(ctrl)
		w := writer.NewMockWriter(ctrl)
		c := commenter.New()
		service := New(log, f, w, c)

		err := service.Add(MetaTemplate("invalid"), language.English, tmp, false)

		assert.Error(t, err)
	})

	t.Run("fetcher error", func(t *testing.T) {
		lang := language.Japanese
		filename := "FUNDING.yml"
		url := fmt.Sprintf(
			"%s/%s/.github/%s",
			source_url.Templates,
			lang,
			filename,
		)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		f := fetcher.NewMockFetcher(ctrl)
		w := writer.NewMockWriter(ctrl)
		c := commenter.New()

		f.EXPECT().Fetch(url).Return(nil, errors.New("fetch error")).Times(1)
		service := New(log, f, w, c)

		err := service.Add(Funding, lang, tmp, false)

		assert.Error(t, err)
	})

	t.Run("writer error", func(t *testing.T) {
		lang := language.Japanese
		filename := "SECURITY.md"
		url := fmt.Sprintf(
			"%s/%s/.github/%s",
			source_url.Templates,
			lang,
			filename,
		)
		data := []byte("template data")
		exceptedPath := filepath.Join(tmp, ".github", filename)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		f := fetcher.NewMockFetcher(ctrl)
		w := writer.NewMockWriter(ctrl)
		c := commenter.New()

		f.EXPECT().Fetch(url).Return(data, nil).Times(1)
		w.EXPECT().
			Write(exceptedPath, c.PrependGeneratedComment(data, format.Markdown, url)).
			Return(errors.New("write error")).
			Times(1)
		service := New(log, f, w, c)

		err := service.Add(Security, lang, tmp, false)

		assert.Error(t, err)
	})
}
