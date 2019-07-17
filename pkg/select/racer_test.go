package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("fastest server wins", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(10 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slow := slowServer.URL
		fast := fastServer.URL

		want := fast
		got, err := Racer(slow, fast)

		if err != nil {
			t.Fatalf("expected no error, but get one '%s'", err)
		}

		if got != want {
			t.Errorf("want '%s', but got '%s'", want, got)
		}
	})

	t.Run("returns error if no server respond in 10 seconds", func(t *testing.T) {
		slowServer := makeDelayedServer(25 * time.Millisecond)
		fastServer := makeDelayedServer(15 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slow := slowServer.URL
		fast := fastServer.URL

		_, err := ConfigurableRacer(slow, fast, 10*time.Millisecond)

		if err == nil {
			t.Fatalf("expected error, but didn't get one")
		}
	})
}

func makeDelayedServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}
