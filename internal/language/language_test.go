package language

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			name  string
			value string
			want  Language
		}{
			{name: "en", value: "en", want: English},
			{name: "ja", value: "ja", want: Japanese},
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
