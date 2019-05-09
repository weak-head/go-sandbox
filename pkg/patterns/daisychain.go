package patterns

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

// BuildDaisyChain is example of building
// and pipelining data thought the chain of
// channels
func BuildDaisyChain(n int) (head, tail chan int) {
	tail = make(chan int)
	head = make(chan int)

	var last, next chan int
	last = head

	for i := 0; i < n; i++ {
		next = make(chan int)
		go f(next, last)
		last = next
	}

	go f(tail, last)

	return
}

// UseDaisyChain is an example of using daisy chain
func UseDaisyChain() {
	head, tail := BuildDaisyChain(200200)

	go func(c chan int) { c <- 0 }(head)

	fmt.Println(<-tail)
}
