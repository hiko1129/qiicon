package qiicon

import (
	"github.com/hiko1129/qiicon/infrastructure"
	"github.com/hiko1129/qiicon/usecase"
)

// Client struct
type Client struct {
	username string
}

// New func
func New(username string) (*Client, error) {
	return &Client{username}, nil
}

// FetchTotalContributionCount func
func (c *Client) FetchTotalContributionCount() (int, error) {
	req := &usecase.FetchTotalContributionCountRequest{Username: c.username}
	client, err := infrastructure.NewContirubutionClient()
	if err != nil {
		return 0, err
	}
	u, err := usecase.NewFetchTotalContributionCount(req, client)
	if err != nil {
		return 0, err
	}
	res, err := u.Exec()
	if err != nil {
		return 0, err
	}

	return res.Contribution, nil
}

// FetchTodayContributionCount func
func (c *Client) FetchTodayContributionCount() (int, error) {
	req := &usecase.FetchTodayContributionCountRequest{Username: c.username}
	client, err := infrastructure.NewContirubutionClient()
	if err != nil {
		return 0, err
	}
	u, err := usecase.NewFetchTodayContributionCount(req, client)
	if err != nil {
		return 0, err
	}
	res, err := u.Exec()
	if err != nil {
		return 0, err
	}

	return res.Contribution, nil
}

// FetchContributions func
func (c *Client) FetchContributions() (map[string]int, error) {
	result := map[string]int{}
	req := &usecase.FetchContributionsRequest{Username: c.username}
	client, err := infrastructure.NewContirubutionClient()
	if err != nil {
		return result, err
	}
	u, err := usecase.NewFetchContributions(req, client)
	if err != nil {
		return result, err
	}
	res, err := u.Exec()
	if err != nil {
		return result, err
	}

	return res.Contributions, nil
}
