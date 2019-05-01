package chans

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func PrintRedits() {
	urls := []string{
		"http://www.reddit.com/r/aww.json",
		"http://www.reddit.com/r/funny.json",
		"http://www.reddit.com/r/programming.json",
	}

	jsonResponses := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()

			// The easy way, but no header:
			// > res, err := http.Get(url)
			// We need header to make reddit work, so...
			client := &http.Client{}
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-agent", "Go bot 0.1")
			res, err := client.Do(req)

			if err != nil {
				log.Fatal(err)
			} else {
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)

				if err != nil {
					log.Fatal(err)
				} else {
					jsonResponses <- string(body)
				}
			}
		}(url)
	}

	go func() {
		for response := range jsonResponses {
			fmt.Println(response)
		}
	}()

	wg.Wait()
	close(jsonResponses)
}
