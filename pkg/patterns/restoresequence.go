package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

// Message is a data from generator with wait signal
type Message struct {
	str  string
	wait chan bool
}

// WaitGenerator is the generator that waits for the signal
// to generate the next value
func WaitGenerator(msg string) <-chan Message {
	c := make(chan Message)

	go func() {
		wait := make(chan bool)
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-wait
		}
	}()

	return c
}

// FanInMsg multiplexes Messages
func FanInMsg(c1, c2, c3 <-chan Message) <-chan Message {
	c := make(chan Message)

	go func() {
		for {
			select {
			case v := <-c1:
				c <- v
			case v := <-c2:
				c <- v
			case v := <-c3:
				c <- v
			}
		}
	}()

	return c
}

// WaitGeneratorInUse example of restoring sequence
func WaitGeneratorInUse() {
	c := FanInMsg(
		WaitGenerator("Howdy"),
		WaitGenerator("Boom Bang"),
		WaitGenerator("Doing Stuff")
	)

	for i := 0; i < 5; i++ {
		v1 := <-c
		v2 := <-c
		v3 := <-c

		fmt.Println(v1.str)
		fmt.Println(v2.str)
		fmt.Println(v3.str)

		v1.wait <- true
		v2.wait <- true
		v3.wait <- true

		fmt.Println()
	}
}
