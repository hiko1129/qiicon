package qiicon_test

import (
	"testing"

	"github.com/hiko1129/qiicon/qiicon"
	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestNew(t *testing.T) {
	_, err := qiicon.New("hiko1129")
	assert.NoError(t, err)
}

func TestFetchTotalContributionCount(t *testing.T) {
	// real
	c, _ := qiicon.New("hiko1129")
	_, err := c.FetchTotalContributionCount()
	assert.NoError(t, err)

	// mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://qiita.com/api/internal/hovercard_users/hiko1129", httpmock.NewStringResponder(200, `{"contribution": 100}`))

	count, err := c.FetchTotalContributionCount()
	assert.Equal(t, 100, count)
	assert.NoError(t, err)

	httpmock.RegisterResponder("GET", "https://qiita.com/api/internal/hovercard_users/-1-1-1-", httpmock.NewStringResponder(404, "not found"))

	c, _ = qiicon.New("-1-1-1-")
	count, err = c.FetchTotalContributionCount()
	assert.Error(t, err)
	assert.Equal(t, 0, count)
}

func TestFetchTodayContributionCount(t *testing.T) {
	// real
	c, _ := qiicon.New("hiko1129")

	_, err := c.FetchTodayContributionCount()
	assert.NoError(t, err)
}

// func TestFetchContributions(t *testing.T) {
// 	// real
// 	c, _ := qiicon.New("hiko1129")

// 	contributions, err := c.FetchContributions()
// 	assert.NotEmpty(t, contributions)
// 	assert.NoError(t, err)
// }
