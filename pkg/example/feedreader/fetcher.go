package feedreader

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rss "github.com/mattn/go-pkg-rss"
)

// An Item is striped-out RSS item
type Item struct {
	Title, GUID string
	Links       []string
}

// A Fetcher fetches Items and returns the time
// when the next fetch should be attempted.
type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

type fakeFetcher struct {
	uri   string
	items []Item
}

type realFetcher struct {
	uri   string
	feed  *rss.Feed
	items []Item
}

// Fetch returns a Fetcher for RSS Items from uri.
func Fetch(uri string, fake bool) Fetcher {
	if fake {
		return fakeFetch(uri)
	}
	return realFetch(uri)
}

func fakeFetch(uri string) Fetcher {
	return &fakeFetcher{uri: uri}
}

func realFetch(uri string) Fetcher {
	ft := &realFetcher{uri: uri}

	newChans := func(feed *rss.Feed, chs []*rss.Channel) {}
	newItems := func(feed *rss.Feed, ch *rss.Channel, items []*rss.Item) {
		for _, item := range items {
			var links []string
			for _, link := range item.Links {
				links = append(links, link.Href)
			}

			ni := Item{Title: item.Title, Links: links}
			if item.Guid != nil {
				ni.GUID = *item.Guid
			}

			ft.items = append(ft.items, ni)
		}
	}

	ft.feed = rss.New(1, true, newChans, newItems)

	return ft
}

func (f *fakeFetcher) Fetch() (items []Item, next time.Time, err error) {
	now := time.Now()
	next = now.Add(time.Duration(rand.Intn(3)) * 700 * time.Millisecond)

	item := Item{
		Title: fmt.Sprintf("Item %d", len(f.items)),
		Links: []string{"Link1", "Link2"},
	}
	item.GUID = strconv.Itoa(rand.Intn(550)) + "_" + item.Title

	f.items = append(f.items, item)
	items = f.items
	return
}

func (f *realFetcher) Fetch() (items []Item, next time.Time, err error) {
	if err = f.feed.Fetch(f.uri, nil); err != nil {
		return
	}

	items = f.items
	f.items = nil

	next = time.Now().Add(time.Duration(f.feed.SecondsTillUpdate()) * time.Second)
	return
}
