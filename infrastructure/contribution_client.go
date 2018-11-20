package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/hiko1129/qiicon/domain/object"
)

const (
	baseEndpoint              = "https://qiita.com/"
	hoverCardUsersAPIEndpoint = "https://qiita.com/api/internal/hovercard_users/"
)

// ContirubutionClient struct
type ContirubutionClient struct {
}

// userActivityChart struct
type userActivityChart struct {
	Data data `json:"data"`
}

type data struct {
	Columns [][]json.Number `json:"columns"`
}

// NewContirubutionClient func
func NewContirubutionClient() (*ContirubutionClient, error) {
	return &ContirubutionClient{}, nil
}

// FetchContributions func
func (c *ContirubutionClient) FetchContributions(username string) (map[string]int, error) {
	contributionURL := baseEndpoint + username + "/" + "contributions"
	con := map[string]int{}

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

		con[date.String()] = int(c)
	}

	return con, nil
}

// FetchTotalContribution func
func (c *ContirubutionClient) FetchTotalContribution(username string) (*object.User, error) {
	res, err := http.Get(hoverCardUsersAPIEndpoint + username)
	var r object.User
	if err != nil {
		return &r, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&r)
	if err != nil {
		return &r, err
	}

	return &r, nil
}
