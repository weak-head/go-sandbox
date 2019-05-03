package main

import (
	"fmt"
	"gobox/pkg/concur/crawler"
)

type FetchProcessor struct {
}

func (fp *FetchProcessor) Process(fi *crawler.FetchedInfo) {
	fmt.Println(fi.Url, fi.Body)
}

func main() {
	crawler.Crawl("http://base.com", 4, crawler.MakeFetcher(), &FetchProcessor{})
}
