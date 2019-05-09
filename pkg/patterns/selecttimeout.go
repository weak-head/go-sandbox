package patterns

import (
	"fmt"
	"time"
)

// SelectTimeoutMessage is an example of using timeouts for
// terminating each message
func SelectTimeoutMessage() {
	c := Generator("Rushing out")

	for i := 0; i < 30; i++ {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("To slow...")
		}
	}
}

// SelectTimeoutConversation is an example of using timeouts for
// terminating the entire conversation
func SelectTimeoutConversation() {
	c := Generator("Wrapping candies")
	t := time.After(3 * time.Second)

Packaging:
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-t:
			fmt.Println("It's time to ship what has been packed...")
			break Packaging
		}
	}
}
