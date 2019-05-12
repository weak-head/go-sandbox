package feedreader

import "fmt"

var (
	feeds = []string{"https://news.ycombinator.com/rss"}
)

func FetchFeeds() {
	fetch := Fetch(feeds[0], false)
	items, _, _ := fetch.Fetch()

	for _, i := range items {
		fmt.Println(i)
	}
}
