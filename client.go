package qiicon

import (
	"encoding/json"
	"net/http"
)

// Client struct
type Client struct {
	username string
}

const (
	hoverCardUsersAPIEndpoint = "http://qiita.com/api/internal/hovercard_users"
)

// HoverCardUserResponse struct
type HoverCardUserResponse struct {
	Contribution int `json:"contribution"`
}

// New func
func New(username string) (*Client, error) {
	return &Client{username}, nil
}

// FetchTotalContributionCount func
func (c *Client) FetchTotalContributionCount() (int, error) {
	res, err := http.Get(hoverCardUsersAPIEndpoint + "/" + c.username)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var r HoverCardUserResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return 0, err
	}

	return r.Contribution, nil
}
