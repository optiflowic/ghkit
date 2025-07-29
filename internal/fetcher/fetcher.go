package fetcher

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./$GOPACKAGE.mock.go
type Fetcher interface {
	Fetch(rawURL string) ([]byte, error)
}
