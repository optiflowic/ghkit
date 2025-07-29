package meta

import (
	"fmt"

	f "github.com/optiflowic/ghkit/internal/format"
	"github.com/optiflowic/ghkit/internal/utils"
)

type MetaTemplate string
type FileInfo struct {
	name   string
	format f.Format
}

const (
	CodeOwners   MetaTemplate = "codeowners"
	Contributing MetaTemplate = "contributing"
	Funding      MetaTemplate = "funding"
	Security     MetaTemplate = "security"
	Support      MetaTemplate = "support"
	All          MetaTemplate = "all"
)

var metaTemplates = map[MetaTemplate]FileInfo{
	CodeOwners: {
		name:   "CODEOWNERS",
		format: f.PlaneText,
	},
	Contributing: {
		name:   "CONTRIBUTING.md",
		format: f.Markdown,
	},
	Funding: {
		name:   "FUNDING.yml",
		format: f.Yaml,
	},
	Security: {
		name:   "SECURITY.md",
		format: f.Markdown,
	},
	Support: {
		name:   "SUPPORT.md",
		format: f.Markdown,
	},
}

var metaTemplateNames = []MetaTemplate{
	CodeOwners,
	Contributing,
	Funding,
	Security,
	Support,
	All,
}

func NewMetaTemplate(value string) (*MetaTemplate, error) {
	template := MetaTemplate(value)
	if !utils.Contains(metaTemplateNames, template) {
		return nil, fmt.Errorf("unsupported template: %s", value)
	}

	return &template, nil
}

func ListAvailable() []string {
	filenames := []string{}
	for _, fileInfo := range metaTemplates {
		filenames = append(filenames, fileInfo.name)
	}

	return filenames
}

func (m MetaTemplate) Get() string {
	return string(m)
}

func find(template MetaTemplate) (*FileInfo, error) {
	fileInfo, ok := metaTemplates[template]
	if !ok {
		return nil, fmt.Errorf("unsupported template: %s", template.Get())
	}
	return &fileInfo, nil
}

func all() []FileInfo {
	fileInfoList := []FileInfo{}
	for _, fileInfo := range metaTemplates {
		fileInfoList = append(fileInfoList, fileInfo)
	}
	return fileInfoList
}
