package chans

import (
	"fmt"
	"sync"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// Sum returns a total sum of an array
func Sum(s []int) int {
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	r1, r2 := <-c, <-c
	return r1 + r2
}

func BufferedChan() {
	c := make(chan int, 100)
	for i := 0; i < 50; i++ {
		c <- i
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func RangeSyncClose() {

	var wg sync.WaitGroup
	wg.Add(2)

	producer := func(c chan int, n int) {
		defer wg.Done()
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	consumer := func(c chan int) {
		defer wg.Done()
		for v := range c {
			fmt.Println(v)
			time.Sleep(400 * time.Millisecond)
		}
	}

	c := make(chan int, 3)
	go producer(c, 10)
	go consumer(c)

	fmt.Println("Waiting to finish...")
	wg.Wait()
	fmt.Println("Finished")
}
