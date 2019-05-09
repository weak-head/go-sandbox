package patterns

import (
	"fmt"
	"time"
)

// SelectTimeout is an example of using timeout for each
// message in select
func SelectTimeout() {
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
