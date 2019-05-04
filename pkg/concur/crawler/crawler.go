package crawler

import (
	"errors"
	"sync"
)

type Fetcher interface {
	Fetch(url string) *FetchedInfo
}

type Processor interface {
	Process(fetched *FetchedInfo)
}

type FetchedInfo struct {
	Url   string
	Body  string
	Links []string
	Er    error
}

var (
	fetchMap = make(map[string]*FetchedInfo)
	mu       sync.RWMutex
)

func Crawl(url string, depth int, fetcher Fetcher, processor Processor) {
	if depth <= 0 {
		return
	}

	var fetched *FetchedInfo

	mu.Lock()
	if _, alreadyFetched := fetchMap[url]; !alreadyFetched {
		fetched = fetcher.Fetch(url)
		fetchMap[url] = fetched
	}
	mu.Unlock()

	if fetched != nil {
		processor.Process(fetched)
		for _, link := range fetched.Links {
			go Crawl(link, depth-1, fetcher, processor)
		}
	}
}

type FakeFetcher struct {
	urls map[string]*FetchedInfo
}

func (fetcher *FakeFetcher) Fetch(url string) *FetchedInfo {
	if info, ok := fetcher.urls[url]; ok {
		return info
	}
	return &FetchedInfo{url, "", nil, errors.New("404 - Not Found")}
}

func MakeFetcher() Fetcher {
	return &FakeFetcher{
		urls: map[string]*FetchedInfo{
			"http://base.com": &FetchedInfo{
				"http://base.com",
				"Base",
				[]string{
					"http://base.com/user",
					"http://base.com/auth",
					"http://base.com/help",
				},
				nil,
			},
		},
	}
}
