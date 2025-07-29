package pr

type PRTemplate string

const (
	PullRequest PRTemplate = "pr"
)

var prTemplates = map[PRTemplate]string{
	PullRequest: "PULL_REQUEST_TEMPLATE.md",
}

func ListAvailable() []string {
	filenames := []string{}
	for _, filename := range prTemplates {
		filenames = append(filenames, filename)
	}

	return filenames
}

func (p PRTemplate) Get() string {
	return string(p)
}
