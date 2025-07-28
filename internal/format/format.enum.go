package format

import "fmt"

type Format string

const (
	Yaml      Format = "yml"
	Markdown  Format = "md"
	PlaneText Format = "txt"
)

func New(value string) (*Format, error) {
	if value != Yaml.Get() && value != Markdown.Get() && value != PlaneText.Get() {
		return nil, fmt.Errorf("unsupported format: %s", value)
	}

	format := Format(value)
	return &format, nil
}

func (f Format) Get() string {
	return string(f)
}
