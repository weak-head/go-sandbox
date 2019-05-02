package counter

import (
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	s.v[key]++
	s.mux.Unlock()
}

func (s *SafeCounter) Val(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}
