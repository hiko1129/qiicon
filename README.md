# qiicon

## Installation

`$ go get github.com/hiko1129/qiicon/qiicon`

## Examples

```Go
import (
	"github.com/hiko1129/qiicon/qiicon"
)

func Example() error {
	c, err := qiicon.New("hiko1129")
	if err != nil {
		return err
	}

	count, err := c.FetchTotalContributionCount()
	if err != nil {
		return err
	}
	fmt.Println(count)

	count, err = c.FetchTodayContributionCount()
	if err != nil {
		return err
	}
	fmt.Println(count)

	contributions, err := c.FetchContributions()
	if err != nil {
		return err
	}
	fmt.Println(contributions)

	return nil
}
```