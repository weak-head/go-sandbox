package main

import (
	pl "gobox/pkg/concur/poller"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://microsoft.com",
		"http://amazon.com",
		"http://ababba.com.com",
		"http://some.other.com",
	}
	pl.Poll(urls)
}
