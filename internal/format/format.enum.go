package format

import "fmt"

type Format string

const (
	Yaml      Format = "yml"
	Markdown  Format = "md"
	PlainText Format = "txt"
)

func New(value string) (*Format, error) {
	if value != Yaml.Get() && value != Markdown.Get() && value != PlainText.Get() {
		return nil, fmt.Errorf("unsupported format: %s", value)
	}

	format := Format(value)
	return &format, nil
}

func (f Format) Get() string {
	return string(f)
}
