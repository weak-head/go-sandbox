package crawler

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

func Crawl(url string, depth int, fetched Fetcher, processor Processor) {
	processor.Process(&FetchedInfo{"abc", "body", nil, nil})
}

type FakeFetcher struct {
	urls map[string]*FetchedInfo
}

func (fetcher *FakeFetcher) Fetch(url string) *FetchedInfo {
	return nil
}

func MakeFetcher() Fetcher {
	return &FakeFetcher{
		urls: map[string]*FetchedInfo{
			"http://base.come": &FetchedInfo{
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
