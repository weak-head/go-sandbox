package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type CountdownOperationsSpy struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) Sleep(d time.Duration) {
	s.durationSlept += d
}

func TestCountdown(t *testing.T) {

	t.Run("count to 3", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		sleeper := &CountdownOperationsSpy{}
		Countdown(sleeper, sleeper)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, sleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, sleeper.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {

	t.Run("sleep once", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

		sleeper.Sleep()

		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})

	t.Run("sleep several times", func(t *testing.T) {
		sleepTime := 5 * time.Second
		sleepCalls := int64(10)
		totalSleep := time.Duration(int64(sleepTime) * sleepCalls)

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

		for i := int64(0); i < sleepCalls; i++ {
			sleeper.Sleep()
		}

		if spyTime.durationSlept != totalSleep {
			t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
