package example

import (
	"fmt"
	"math/rand"
	"time"
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
