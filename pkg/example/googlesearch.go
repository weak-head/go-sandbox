package example

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
	All   = []Search{Web, Image, Video}
)

// Result is a search result
type Result string

// Search is a search delegate
type Search func(query string) Result

// fakeSearch creates a fake search of the specified kind
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %s", kind, query))
	}
}

// Google1 -> sequential search one by one
func Google1(query string) (results []Result) {
	for _, s := range All {
		results = append(results, s(query))
	}
	return
}

// DoSearch starts the search and prints out the results
func DoSearch() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	results := Google1("golang")
	elapsed := time.Since(start)

	// fmt.Printf("%#v\n", results)
	var strs []string
	for _, r := range results {
		strs = append(strs, string(r))
	}
	fmt.Println(strings.Join(strs, "\n"))
	fmt.Println(elapsed)
}
