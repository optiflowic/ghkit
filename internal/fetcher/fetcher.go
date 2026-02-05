package fetcher

import "context"

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./$GOPACKAGE.mock.go
type Fetcher interface {
	Fetch(ctx context.Context, rawURL string) ([]byte, error)
}
