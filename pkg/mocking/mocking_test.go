package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpySleeper struct {
	Calls int
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("count to 3", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

		if sleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", sleeper.Calls)
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
