package feedreader

import (
	"fmt"
	"time"
)

var (
	feeds = []string{
		"https://news.ycombinator.com/rss",
		"https://www.allthingsdistributed.com/index.xml",
		"https://www.joelonsoftware.com/feed",
	}
)

func FetchFeeds() {
	var subs []Subscription
	for _, feed := range feeds {
		subs = append(subs, Subscribe(Fetch(feed, false)))
	}
	merged := Merge(subs...)

	time.AfterFunc(10*time.Second, func() {
		fmt.Println("Closed:", merged.Close())
	})

	for it := range merged.Updates() {
		fmt.Println(it)
	}
}
