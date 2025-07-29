package pr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListAvailable(t *testing.T) {
	list := ListAvailable()

	assert.Contains(t, list, "PULL_REQUEST_TEMPLATE.md")
}

func Test_Get(t *testing.T) {
	got := PullRequest.Get()

	assert.Equal(t, "pr", got)
}
