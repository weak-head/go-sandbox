package patterns

import (
	"fmt"
	"math/rand"
)

func cleanup() {
	fmt.Println("Cleaning up")
}

// QuitGenerator stops generating new items
// when requested
func QuitGenerator(msg string, quit chan bool) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-quit:
				cleanup()
				quit <- true
				return
			}
		}
	}()

	return c
}

// QuitUsage example of quiting the channel
// with cleanup
func QuitUsage() {
	q := make(chan bool)
	c := QuitGenerator("Unwrapping unicorns", q)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}

	q <- true
	fmt.Println("Waiting for clean up")
	<-q
	fmt.Println("Done!")
}
