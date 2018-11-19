package qiicon

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hiko1129/qiicon/domain"
)

// Client struct
type Client struct {
	username string
}

const (
	hoverCardUsersAPIEndpoint = "https://qiita.com/api/internal/hovercard_users/"
)

// HoverCardUserResponse struct
type hoverCardUserResponse struct {
	Contribution int `json:"contribution"`
}

// New func
func New(username string) (*Client, error) {
	return &Client{username}, nil
}

// FetchTotalContributionCount func
func (c *Client) FetchTotalContributionCount() (int, error) {
	res, err := http.Get(hoverCardUsersAPIEndpoint + c.username)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var r hoverCardUserResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return 0, err
	}

	return r.Contribution, nil
}

// FetchTodayContributionCount func
func (c *Client) FetchTodayContributionCount() (int64, error) {
	contributions, err := c.extractContributions()
	if err != nil {
		return 0, err
	}

	t := time.Now()
	return contributions[t.Format("2006-01-02")], nil
}

// FetchContributions func
func (c *Client) FetchContributions() (map[string]int64, error) {
	return c.extractContributions()
}

func (c *Client) extractContributions() (map[string]int64, error) {
	e, err := domain.NewExtractor(c.username)
	if err != nil {
		return map[string]int64{}, err
	}

	return e.ExtractContributions()
}
