package context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := r.Context()
		out := make(chan string, 1)

		go func() {
			out <- store.Fetch()
		}()

		select {
		case fetched := <-out:
			fmt.Fprintf(w, fetched)
		case <-context.Done():
			store.Cancel()
		}
	}
}
