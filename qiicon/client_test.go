package qiicon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchTotalContributionCount(t *testing.T) {
	c, err := New("hiko1129")
	assert.NoError(t, err)

	_, err = c.FetchTotalContributionCount()
	assert.NoError(t, err)

	c, _ = New("-1-1-1-")
	count, err := c.FetchTotalContributionCount()
	assert.Error(t, err)
	assert.Equal(t, 0, count)
}
