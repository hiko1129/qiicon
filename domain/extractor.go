package domain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseEndpoint = "https://qiita.com/"
)

// Extractor struct
type Extractor struct {
	username string
}

type userActivityChart struct {
	Data data `json:"data"`
}

type data struct {
	Columns [][]json.Number `json:"columns"`
}

// NewExtractor func
func NewExtractor(username string) (*Extractor, error) {
	return &Extractor{username}, nil
}

// ExtractContributions func
func (e *Extractor) ExtractContributions() (map[string]int64, error) {
	contributionURL := baseEndpoint + e.username + "/" + "contributions"
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
