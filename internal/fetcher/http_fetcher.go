package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/optiflowic/ghkit/internal/logger"
)

type HttpFetcher struct {
	log logger.Logger
}

func New(log logger.Logger) HttpFetcher {
	return HttpFetcher{log: log}
}

func (f HttpFetcher) Fetch(rawURL string) ([]byte, error) {
	f.log.Debug("Starting fetch", "url", rawURL)

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		f.log.Error("Invalid URL", "url", rawURL, "error", err)
		return nil, fmt.Errorf("invalid url: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		f.log.Error("Failed to create HTTP request", "error", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
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
