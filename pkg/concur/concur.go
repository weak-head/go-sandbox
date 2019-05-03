package concur

import (
	"fmt"
	"gobox/pkg/concur/chans"
	"gobox/pkg/concur/counter"
	"gobox/pkg/concur/routine"
	"gobox/pkg/concur/tree"
	"time"
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

func TestCounter() {
	cnt1 := counter.New()

	for i := 0; i < 1000; i++ {
		go func(i int) {
			if i%2 == 0 {
				cnt1.Inc("abc")
			} else {
				cnt1.Dec("abc")
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Printf("Val %d\n", cnt1.Val("abc"))

	// this will cause errors
	for i := 0; i < 1000; i++ {
		go func(i int) {
			// and recovery will not help
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f")
				}
			}()

			if i%2 == 0 {
				cnt1.UnsafeInc("abc")
			} else {
				cnt1.UnsafeDec("abc")
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Printf("Val %d\n", cnt1.Val("abc"))
}
