package language

import "fmt"

type Language string

const (
	English  Language = "en"
	Japanese Language = "ja"
)

func New(value string) (*Language, error) {
	if value != English.Get() && value != Japanese.Get() {
		return nil, fmt.Errorf("unsupported language: %s", value)
	}

	l := Language(value)
	return &l, nil
}

func (l Language) Get() string {
	return string(l)
}
