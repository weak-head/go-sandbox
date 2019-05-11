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

// Google2 -> run concurrently and wait for results
func Google2(query string) (results []Result) {
	c := make(chan Result)
	for _, s := range All {
		go func(c chan Result, s Search, q string) {
			c <- s(q)
		}(c, s, query)
	}

	for i := 0; i < len(All); i++ {
		results = append(results, <-c)
	}
	return
}

// DoSearch starts the search and prints out the results
func DoSearch(search func(query string) []Result) {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	results := search("golang")
	elapsed := time.Since(start)

	// fmt.Printf("%#v\n", results)
	var strs []string
	for _, r := range results {
		strs = append(strs, string(r))
	}
	fmt.Println(strings.Join(strs, "\n"))
	fmt.Println(elapsed)
}

// RunSearches runs all the search cases
func RunSearches() {
	fmt.Println("\nGoogle1")
	DoSearch(Google1)

	fmt.Println("\nGoogle2")
	DoSearch(Google2)
}
