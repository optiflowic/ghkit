package meta

import (
	"testing"

	"github.com/optiflowic/ghkit/internal/format"
	"github.com/stretchr/testify/assert"
)

func Test_NewMetaTemplate(t *testing.T) {
	t.Run("valid template", func(t *testing.T) {
		tests := []struct {
			name  string
			value string
		}{
			{name: "codeowners", value: "codeowners"},
			{name: "contributing", value: "contributing"},
			{name: "funding", value: "funding"},
			{name: "security", value: "security"},
			{name: "support", value: "support"},
			{name: "all", value: "all"},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				template, err := NewMetaTemplate(tt.value)

				assert.Equal(t, MetaTemplate(tt.value), *template)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("invalid template", func(t *testing.T) {
		template, err := NewMetaTemplate("invalid")

		assert.Nil(t, template)
		assert.Error(t, err)
	})
}

func Test_ListAvailable(t *testing.T) {
	list := ListAvailable()

	assert.Contains(t, list, "CODEOWNERS")
	assert.Contains(t, list, "CONTRIBUTING.md")
	assert.Contains(t, list, "FUNDING.yml")
	assert.Contains(t, list, "SECURITY.md")
	assert.Contains(t, list, "SUPPORT.md")
}

func Test_Get(t *testing.T) {
	got := CodeOwners.Get()

	assert.Equal(t, "codeowners", got)
}

func Test_find(t *testing.T) {
	t.Run("found file", func(t *testing.T) {
		got, err := find(CodeOwners)

		assert.Equal(t, FileInfo{
			name:   "CODEOWNERS",
			format: format.PlaneText,
		}, *got)
		assert.NoError(t, err)
	})

	t.Run("unsupported template", func(t *testing.T) {
		template := MetaTemplate("invalid")

		got, err := find(template)

		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func Test_all(t *testing.T) {
	got := all()

	assert.Contains(t, got, FileInfo{
		name:   "CODEOWNERS",
		format: format.PlaneText,
	})
	assert.Contains(t, got, FileInfo{
		name:   "CONTRIBUTING.md",
		format: format.Markdown,
	})
	assert.Contains(t, got, FileInfo{
		name:   "FUNDING.yml",
		format: format.Yaml,
	})
	assert.Contains(t, got, FileInfo{
		name:   "SECURITY.md",
		format: format.Markdown,
	})
	assert.Contains(t, got, FileInfo{
		name:   "SUPPORT.md",
		format: format.Markdown,
	})
}
