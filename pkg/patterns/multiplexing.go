package patterns

import (
	"fmt"
)

// FanIn dirty way to multiplex two channels
func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

// Usage is an example if using simple channel mutliplexing
func FanInUsage() {
	c := FanIn(Generator("Hey"), Generator("Wow Wow"))
	for i := 0; i < 10; i++ {
		fmt.Printf("Received: %s\n", <-c)
	}
	fmt.Printf("Done\n")
}
