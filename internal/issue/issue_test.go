package issue

import (
	"testing"

	"github.com/optiflowic/ghkit/internal/format"
	"github.com/stretchr/testify/assert"
)

func Test_NewIssueTemplate(t *testing.T) {
	t.Run("valid template", func(t *testing.T) {
		tests := []struct {
			name  string
			value string
		}{
			{name: "bug", value: "bug"},
			{name: "feature", value: "feature"},
			{name: "question", value: "question"},
			{name: "task", value: "task"},
			{name: "docs", value: "docs"},
			{name: "feedback", value: "feedback"},
			{name: "config", value: "config"},
			{name: "all", value: "all"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				template, err := NewIssueTemplate(tt.value)

				assert.Equal(t, IssueTemplate(tt.value), *template)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("invalid template", func(t *testing.T) {
		template, err := NewIssueTemplate("invalid")

		assert.Nil(t, template)
		assert.Error(t, err)
	})
}

func Test_ListAvailable(t *testing.T) {
	list := ListAvailable()

	assert.Contains(t, list, "bug.yml")
	assert.Contains(t, list, "bug.md")
	assert.Contains(t, list, "feature.yml")
	assert.Contains(t, list, "feature.md")
	assert.Contains(t, list, "question.yml")
	assert.Contains(t, list, "question.md")
	assert.Contains(t, list, "task.yml")
	assert.Contains(t, list, "task.md")
	assert.Contains(t, list, "docs.yml")
	assert.Contains(t, list, "docs.md")
	assert.Contains(t, list, "feedback.yml")
	assert.Contains(t, list, "feedback.md")
	assert.Contains(t, list, "config.yml")
}

func Test_Get(t *testing.T) {
	got := Bug.Get()

	assert.Equal(t, "bug", got)
}

func Test_find(t *testing.T) {
	t.Run("found file", func(t *testing.T) {
		got, err := find(Bug, format.Markdown)

		assert.Equal(t, "bug.md", *got)
		assert.NoError(t, err)
	})

	t.Run("unsupported template", func(t *testing.T) {
		unknown := IssueTemplate("unknown")

		got, err := find(unknown, format.Yaml)

		assert.Nil(t, got)
		assert.Error(t, err)
	})

	t.Run("unsupported format", func(t *testing.T) {
		got, err := find(Config, format.Markdown)

		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func Test_all(t *testing.T) {
	t.Run("valid format", func(t *testing.T) {
		got, err := all(format.Markdown)

		assert.Contains(t, got, "bug.md")
		assert.Contains(t, got, "feature.md")
		assert.Contains(t, got, "question.md")
		assert.Contains(t, got, "task.md")
		assert.Contains(t, got, "docs.md")
		assert.Contains(t, got, "feedback.md")
		assert.Contains(t, got, "config.yml")
		assert.NoError(t, err)
	})

	t.Run("invalid format", func(t *testing.T) {
		invalid := format.Format("unknown")

		got, err := all(invalid)

		assert.Nil(t, got)
		assert.Error(t, err)
	})
}
