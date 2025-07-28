package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			name  string
			value string
			want  Format
		}{
			{name: "yaml", value: "yml", want: Yaml},
			{name: "markdown", value: "md", want: Markdown},
			{name: "planeText", value: "txt", want: PlaneText},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := New(tt.value)

				assert.Equal(t, tt.want, *got)
				assert.NoError(t, err)
			})
		}
	})

	t.Run("error", func(t *testing.T) {
		got, err := New("invalid")

		assert.Nil(t, got)
		assert.Error(t, err)
	})
}

func Test_Get(t *testing.T) {
	value := "yml"
	format, err := New(value)
	assert.NoError(t, err)

	got := format.Get()

	assert.Equal(t, value, got)
}
