package counter

import (
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func New() *SafeCounter {
	return &SafeCounter{v: make(map[string]int)}
}

func (s *SafeCounter) UnsafeInc(key string) {
	s.v[key]++
}

func (s *SafeCounter) UnsafeDec(key string) {
	s.v[key]--
}

func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	s.v[key]++
	s.mux.Unlock()
}

func (s *SafeCounter) Dec(key string) {
	s.mux.Lock()
	s.v[key]--
	s.mux.Unlock()
}

func (s *SafeCounter) Val(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}
