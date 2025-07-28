package fetcher

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/optiflowic/ghkit/internal/logger"
	"github.com/stretchr/testify/assert"
)

func Test_Fetch(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res := "hello world"
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprint(w, res)
		}))
		defer server.Close()

		f := New(logger.NewFromFlags(false, false))
		data, err := f.Fetch(server.URL)

		assert.Equal(t, res, string(data))
		assert.NoError(t, err)
	})

	t.Run("client error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "forbidden", http.StatusForbidden)
		}))
		defer server.Close()

		f := New(logger.NewFromFlags(false, false))
		data, err := f.Fetch(server.URL)

		assert.Nil(t, data)
		assert.Error(t, err)
	})

	t.Run("connection error", func(t *testing.T) {
		f := New(logger.NewFromFlags(false, false))
		data, err := f.Fetch("http://127.0.0.1:0")

		assert.Nil(t, data)
		assert.Error(t, err)
	})
}
