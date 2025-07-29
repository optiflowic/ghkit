package writer

//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=./$GOPACKAGE.mock.go
type Writer interface {
	Write(path string, data []byte) error
}
