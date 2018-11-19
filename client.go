package qiicon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Client struct
type Client struct {
	username string
}

type userActivityChart struct {
	Data data `json:"data"`
}

type data struct {
	Columns [][]json.Number `json:"columns"`
}

const (
	baseEndpoint              = "https://qiita.com/"
	hoverCardUsersAPIEndpoint = baseEndpoint + "api/internal/hovercard_users/"
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
	res, err := http.Get(hoverCardUsersAPIEndpoint + c.username)
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
	contributionURL := baseEndpoint + c.username + "/" + "contributions"
	con := map[string]int64{}

	res, err := http.Get(contributionURL)
	if err != nil {
		return con, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return con, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return con, err
	}

	jsonStr, _ := doc.Find(".userActivityChart").Attr("data-props")
	fmt.Println(jsonStr)
	jsonBytes := ([]byte)(jsonStr)

	var u userActivityChart
	err = json.Unmarshal(jsonBytes, &u)
	if err != nil {
		return con, err
	}

	dates := u.Data.Columns[0][1:]
	contributions := u.Data.Columns[1][1:]

	for i, date := range dates {
		c, err := contributions[i].Int64()
		if err != nil {
			return con, err
		}

		con[date.String()] = c
	}

	return con, nil
}
