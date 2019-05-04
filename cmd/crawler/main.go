package main

import (
	"fmt"
	"gobox/pkg/concur/crawler"
	"time"
)

type FetchProcessor struct {
}

func (fp *FetchProcessor) Process(fi *crawler.FetchedInfo) {
	if fi.Er != nil {
		fmt.Printf("[%s] -> %s\n", fi.Url, fi.Er)
	} else {
		fmt.Printf("[%s]: %s\n", fi.Url, fi.Body)
	}
}

func main() {
	crawler.Crawl("http://base.com", 4, crawler.MakeFetcher(), &FetchProcessor{})
	time.Sleep(time.Second)
	fmt.Println("exit")
}
