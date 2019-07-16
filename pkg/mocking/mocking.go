package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}

	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
