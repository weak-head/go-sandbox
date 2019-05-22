package feedreader

import "time"

// A Subscription delivers items over a channel.
type Subscription interface {
	Updates() <-chan Item
	Close() error
}

// Subscribe returns a new subscription that uses
// Fetcher to fetch items.
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
		closing: make(chan chan error),
	}
	go s.loop()
	return s
}

// sub implements the Subscription interface
type sub struct {
	fetcher Fetcher         // fetches items
	updates chan Item       // sends fetched items
	closing chan chan error // close fetch channel
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	ec := make(chan error)
	s.closing <- ec
	return <-ec
}

// the main subscription loop
func (s *sub) loop() {

	const maxPending = 5

	type fetchResult struct {
		fetched []Item
		next    time.Time
		err     error
	}

	var (
		fetchDone chan fetchResult        // when fetch is running this is not nil
		pending   []Item                  // new feed items to report
		next      time.Time               // next fetch time
		seen      = make(map[string]bool) // already reported feeds
		err       error                   // last error
	)

	for {

		// Evaluate the current wait delay before the next fetch
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}

		// In case if fetch is not running wait for the delay and
		// trigger the next fetch via the startFetch cannel
		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending {
			startFetch = time.After(fetchDelay)
		}

		// If there are some pending items we need to send them
		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}

		select {
		// Close the subscription, clean up the resources
		// and terminate the loop
		case ec := <-s.closing:
			close(s.updates)
			ec <- err
			return

			// At least one
		}
	}
}
