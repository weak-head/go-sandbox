package concur

import (
	"fmt"
	"gobox/pkg/concur/chans"
	"gobox/pkg/concur/routine"
	"gobox/pkg/concur/tree"
)

func DoConcur() {
	routine.DoSay(5)
}

func RunChans() {
	a := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(chans.Sum(a))

	chans.BufferedChan()

	chans.RangeSyncClose()

	chans.SelectMany()

	chans.SelectDefault()
}

func RunRedits() {
	chans.PrintRedits()
}

func TestSame() {
	for n := 80; n < 100; n++ {
		a := tree.New(n)
		b := tree.New(n)

		same := tree.Same(a, b)
		if !same {
			panic(fmt.Sprintf("Not same for %d\n", n))
		} else {
			fmt.Printf("Same for %d\n", n)
		}
	}

	for n := 80; n < 100; n++ {
		a := tree.New(n)
		b := tree.New(n + 1)

		same := tree.Same(a, b)
		if same {
			panic(fmt.Sprintf("Same for %d %d\n", n, n+1))
		} else {
			fmt.Printf("Not same for %d and %d\n", n, n+1)
		}
	}
}
