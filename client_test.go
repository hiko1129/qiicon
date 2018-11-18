package qiicon_test

import (
	"testing"

	"github.com/hiko1129/qiicon"
	"github.com/stretchr/testify/assert"
)

func TestFetchTotalContributionCount(t *testing.T) {
	c, err := qiicon.New("hiko1129")
	assert.NoError(t, err)

	_, err = c.FetchTotalContributionCount()
	assert.NoError(t, err)

	c, _ = qiicon.New("-1-1-1-")
	count, err := c.FetchTotalContributionCount()
	assert.Error(t, err)
	assert.Equal(t, 0, count)
}
