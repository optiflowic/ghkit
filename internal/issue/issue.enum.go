package issue

import (
	"fmt"

	f "github.com/optiflowic/ghkit/internal/format"
	"github.com/optiflowic/ghkit/internal/utils"
)

type IssueTemplate string
type FileMap map[f.Format]string

const (
	Bug      IssueTemplate = "bug"
	Feature  IssueTemplate = "feature"
	Question IssueTemplate = "question"
	Task     IssueTemplate = "task"
	Docs     IssueTemplate = "docs"
	Feedback IssueTemplate = "feedback"
	Config   IssueTemplate = "config"
	All      IssueTemplate = "all"
)

var issueTemplates = map[IssueTemplate]FileMap{
	Bug: {
		f.Yaml:     "bug.yml",
		f.Markdown: "bug.md",
	},
	Feature: {
		f.Yaml:     "feature.yml",
		f.Markdown: "feature.md",
	},
	Question: {
		f.Yaml:     "question.yml",
		f.Markdown: "question.md",
	},
	Task: {
		f.Yaml:     "task.yml",
		f.Markdown: "task.md",
	},
	Docs: {
		f.Yaml:     "docs.yml",
		f.Markdown: "docs.md",
	},
	Feedback: {
		f.Yaml:     "feedback.yml",
		f.Markdown: "feedback.md",
	},
	Config: {
		f.Yaml: "config.yml",
	},
}

var issueTemplateNames = []IssueTemplate{
	Bug,
	Feature,
	Question,
	Task,
	Docs,
	Feedback,
	Config,
	All,
}

func NewIssueTemplate(value string) (*IssueTemplate, error) {
	template := IssueTemplate(value)
	if !utils.Contains(issueTemplateNames, template) {
		return nil, fmt.Errorf("unsupported template: %s", value)
	}

	return &template, nil
}

func ListAvailable() []string {
	filenames := []string{}
	for issueTemplate, fileMap := range issueTemplates {
		if issueTemplate == Config {
			filenames = append(filenames, fileMap[f.Yaml])
			continue
		}

		filenames = append(filenames, fileMap[f.Markdown], fileMap[f.Yaml])
	}

	return filenames
}

func (i IssueTemplate) Get() string {
	return string(i)
}

func find(template IssueTemplate, format f.Format) (*string, error) {
	fileMap, ok := issueTemplates[template]
	if !ok {
		return nil, fmt.Errorf("unsupported template: %s", template.Get())
	}

	filename, ok := fileMap[format]
	if !ok {
		return nil, fmt.Errorf("unsupported format: %s", format.Get())
	}

	return &filename, nil
}

func all(format f.Format) ([]string, error) {
	filenames := []string{}
	for issueTemplate, fileMap := range issueTemplates {
		if issueTemplate == Config {
			filenames = append(filenames, fileMap[f.Yaml])
			continue
		}

		filename, ok := fileMap[format]
		if !ok {
			return nil, fmt.Errorf("unsupported format: %s", format.Get())
		}

		filenames = append(filenames, filename)
	}

	return filenames, nil
}
