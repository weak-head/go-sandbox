package routine

import (
	"fmt"
	"time"
)

func say(s string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Print(s)
	}
}

// DoSay - basic example of goroutine
func DoSay(n int) {
	go say("hello ", 5)
	go say("brave new ", 5)
	say("world\n", 5)
}
