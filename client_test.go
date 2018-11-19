package qiicon_test

import (
	"testing"

	"github.com/hiko1129/qiicon"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_, err := qiicon.New("hiko1129")
	assert.NoError(t, err)
}

func TestFetchTotalContributionCount(t *testing.T) {
	c, _ := qiicon.New("hiko1129")

	_, err := c.FetchTotalContributionCount()
	assert.NoError(t, err)

	c, _ = qiicon.New("-1-1-1-")
	count, err := c.FetchTotalContributionCount()
	assert.Error(t, err)
	assert.Equal(t, 0, count)
}

func TestFetchTodayContributionCount(t *testing.T) {
	c, _ := qiicon.New("hiko1129")

	_, err := c.FetchTodayContributionCount()
	assert.NoError(t, err)
}

func TestFetchContributions(t *testing.T) {
	c, _ := qiicon.New("hiko1129")

	contributions, err := c.FetchContributions()
	assert.NotEmpty(t, contributions)
	assert.NoError(t, err)
}
