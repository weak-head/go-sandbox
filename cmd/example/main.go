package main

import (
	fr "gobox/pkg/example/feedreader"
	ex "gobox/pkg/example/search"
)

func main() {
	ex.RunSearches()
	fr.FetchFeeds()
}
