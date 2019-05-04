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

		var wg sync.WaitGroup
		wg.Add(len(fetched.Links))

		for _, link := range fetched.Links {
			go func(link string) {
				defer wg.Done()
				Crawl(link, depth-1, fetcher, processor)
			}(link)
		}

		wg.Wait()
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
				"Root",
				[]string{
					"http://base.com/user",
					"http://base.com/auth",
					"http://base.com/help",
				},
				nil,
			},
			"http://base.com/user": &FetchedInfo{
				"http://base.com/user",
				"User",
				[]string{
					"http://base.com",
					"http://base.com/auth",
					"http://base.com/signin",
					"http://base.com/signup",
				},
				nil,
			},
			"http://base.com/signup": &FetchedInfo{
				"http://base.com/signup",
				"SignUp",
				[]string{
					"http://base.com",
					"http://base.com/user",
					"http://base.com/signin",
					"http://base.com/auth",
					"http://base.com/recover",
				},
				nil,
			},
			"http://base.com/recover": &FetchedInfo{
				"http://base.com/recover",
				"Recover",
				[]string{
					"http://base.com/help",
					"http://base.com/schedulecall",
					"http://base.com/auth",
				},
				nil,
			},
			"http://base.com/schedulecall": &FetchedInfo{
				"http://base.com/schedulecall",
				"ScheduleCall",
				[]string{
					"http://base.com",
					"http://base.com/auth",
					"http://base.com/news",
				},
				nil,
			},
		},
	}
}
