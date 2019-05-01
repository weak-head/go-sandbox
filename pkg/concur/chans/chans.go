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

func SelectMany() {
	var wg sync.WaitGroup
	wg.Add(6)

	c := make(chan int, 3)
	p := make(chan int, 3)
	r := make(chan int)

	// 3 producers
	for i := 0; i < 3; i++ {
		go func(ch chan int, n int) {
			defer wg.Done()
			defer fmt.Printf("producer c%d is done\n", n)
			for j := 0; j < 50; j++ {
				ch <- j
			}
		}(c, i)
	}

	// another 3 producers
	for i := 0; i < 3; i++ {
		go func(ch chan int, n int) {
			defer wg.Done()
			defer fmt.Printf("producer p%d is done\n", n)
			for j := 0; j < 50; j++ {
				ch <- j
			}
		}(p, i)
	}

	// 1 consumer
	go func(ch1, ch2 chan int) {
		s := 0
		nwait := 3
		for {
			select {
			case v1 := <-ch1:
				s += v1
			case v2 := <-ch2:
				s += v2
			case <-time.After(1 * time.Second):
				if nwait <= 0 {
					fmt.Println("Exit consumer")
					r <- s
					return
				}
				fmt.Println("No input from producers")
				nwait--
			}
		}
	}(c, p)

	fmt.Println("Waiting for producers to finish")
	wg.Wait()
	fmt.Println("All producers are finished")
	fmt.Printf("Result %d\n", <-r)
	close(c)
	close(p)
}
