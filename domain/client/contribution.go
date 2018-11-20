package client

import "github.com/hiko1129/qiicon/domain/object"

// Contribution interface
type Contribution interface {
	FetchContributions(string) (map[string]int, error)
	FetchTotalContribution(string) (*object.User, error)
}
