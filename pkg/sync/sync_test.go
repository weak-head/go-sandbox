package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

	t.Run("increment the counter by 3 leaves it at 3", func(t *testing.T) {
		counter := NewCounter()

		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(wg *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	val := got.Value()
	if val != want {
		t.Errorf("Expected to have %d, but got %d", want, got)
	}
}
