package poller

type (
	Resource struct {
		url    string
		errors int
	}

	State struct {
		url    string
		status string
	}

	HttpPoller struct{}
	Poller     interface {
		Poll(url string) (status string)
	}
)

const (
	maxRetry = 3
)

func PollUrl(poller Poller, in <-chan *Resource, out chan<- *Resource, state chan<- *State) {
	for i := range in {
		status := poller.Poll(i.url)
		state <- &State{i.url, status}

		if status == "error" {
			i.errors++
		}
		out <- i
	}
}

func Poll(urls []string) map[string]string {

}
