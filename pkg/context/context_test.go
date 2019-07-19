package context

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-time.After(10 * time.Millisecond):
				result += string(c)
			case <-ctx.Done():
				s.t.Log("Canceled store")
				return
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyWriter struct {
	written bool
}

func (s *SpyWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestHandler(t *testing.T) {
	data := "hello!"

	t.Run("server returns the fetched content", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got '%s' want '%s'", response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		writer := &SpyWriter{}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		svr.ServeHTTP(writer, request)

		if writer.written {
			t.Errorf("a response should not have been written")
		}
	})
}
