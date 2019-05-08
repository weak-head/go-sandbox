package pattern

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator is a functions that returns a
// receive-only channel
func Generator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("[%02d] -> %s", i, msg)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// Usage is an example if using Generator pattern
func Usage() {
	c := Generator("message")
	for i := 0; i < 10; i++ {
		fmt.Printf("Received: %s\n", <-c)
	}
	fmt.Printf("Done\n")
}
