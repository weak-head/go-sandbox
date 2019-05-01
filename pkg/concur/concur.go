package concur

import (
	"fmt"
	"gobox/pkg/concur/chans"
	"gobox/pkg/concur/routine"
)

func DoConcur() {
	routine.DoSay(5)
}

func RunChans() {
	a := []int{10, 20, 30, 40, 50, 60}
	fmt.Println(chans.Sum(a))

	chans.BufferedChan()

	chans.RangeSyncClose()
}

func RunRedits() {
	chans.PrintRedits()
}
