package usecase

import (
	"time"

	"github.com/hiko1129/qiicon/domain/client"
	"github.com/pkg/errors"
)

// FetchTodayContributionCountRequest struct
type FetchTodayContributionCountRequest struct {
	Username string
}

// FetchTodayContributionCountResponse struct
type FetchTodayContributionCountResponse struct {
	Contribution int
}

// FetchTodayContributionCount struct
type FetchTodayContributionCount struct {
	request *FetchTodayContributionCountRequest
	client  client.Contribution
}

// NewFetchTodayContributionCount func
func NewFetchTodayContributionCount(request *FetchTodayContributionCountRequest, client client.Contribution) (*FetchTodayContributionCount, error) {
	return &FetchTodayContributionCount{request: request, client: client}, nil
}

// Exec func
func (f *FetchTodayContributionCount) Exec() (*FetchTodayContributionCountResponse, error) {
	fe := &FetchTodayContributionCountResponse{}
	contributions, err := f.client.FetchContributions(f.request.Username)
	if err != nil {
		return fe, errors.Wrap(err, "fetch today contribution failed")
	}

	t := time.Now()
	fe.Contribution = contributions[t.Format("2006-01-02")]

	return fe, nil
}
