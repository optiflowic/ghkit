package fetcher

import (
	"fmt"
	"io"
	"net/http"

	"github.com/optiflowic/ghkit/internal/logger"
)

type Fetcher struct {
	log logger.Logger
}

func New(log logger.Logger) Fetcher {
	return Fetcher{log: log}
}

func (f Fetcher) Fetch(url string) ([]byte, error) {
	f.log.Debug("Starting fetch", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		f.log.Error("HTTP request failed", "error", err)
		return nil, fmt.Errorf("failed to fetch: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			f.log.Error("Failed to close response body", "error", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		f.log.Error("Unexpected HTTP status", "status", resp.Status)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		f.log.Error("Failed to read body", "error", err)
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	f.log.Debug("Fetch successful", "bytes", len(body))
	return body, nil
}
