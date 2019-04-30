package chans

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
