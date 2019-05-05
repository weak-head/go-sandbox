package poller

import (
	"log"
	"net/http"
	"time"
)

type (
	Resource struct {
		url    string
		errors int
		status string
		code   int
	}

	State struct {
		url    string
		status string
	}

	HttpPoller struct{}
	Poller     interface {
		Poll(url string) (code int, status string)
	}
)

const (
	maxRetry = 3
)

func (hp *HttpPoller) Poll(url string) (code int, status string) {
	resp, err := http.Head(url)

	if err != nil {
		log.Println("Error", url, err)
		return 0, err.Error()
	}

	return resp.StatusCode, resp.Status
}

func mkStateMonitor(pollTime time.Duration) chan<- *State {
	state := make(chan *State)
	stateMap := make(map[string]string)
	timer := time.NewTicker(pollTime)

	go func() {
		for {
			select {
			case <-timer.C:
				logState(stateMap)
			case s := <-state:
				stateMap[s.url] = s.status
			}
		}
	}()

	return state
}

func logState(stateMap map[string]string) {
	for k, v := range stateMap {
		log.Printf("%s -> %s\n", k, v)
	}
}

func PollUrl(poller Poller, in <-chan *Resource, out chan<- *Resource, state chan<- *State) {
	for i := range in {
		code, status := poller.Poll(i.url)
		state <- &State{i.url, status}

		// We are interested only in 200 OK
		if code != 200 {
			i.errors++
		} else {
			i.status = status
			i.code = code
		}
		out <- i
	}
}

func RetryFailed(done <-chan *Resource, retry chan<- *Resource) {
	for r := range done {
		// OK
		if r.code == 200 {
			continue
		}

		// Failed
		if r.errors < maxRetry {
			go func(r *Resource) {
				time.Sleep(1 * time.Second)
				retry <- r
			}(r)
		}
	}
}

func Poll(urls []string) map[string]string {
	return nil
}
