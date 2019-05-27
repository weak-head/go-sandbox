package feedreader

// Implements Subscription interface
// and can aggregate multiple subscriptions
// into one. Similar to Composite.
type merge struct {
	subs    []Subscription
	updates chan Item
	quit    chan struct{}
	err     chan error
}

func (m *merge) Close() (err error) {
	close(m.quit)
	for range m.subs {
		if e := <-m.err; e != nil {
			err = e
		}
	}
	close(m.updates)
	return
}

func (m *merge) Updates() <-chan Item {
	return m.updates
}

// Merge follows the Fan-In pattern and merges
// multiple Subscriptions into one
func Merge(subs ...Subscription) Subscription {
	m := &merge{
		subs:    subs,
		updates: make(chan Item),
		quit:    make(chan struct{}),
		err:     make(chan error),
	}

	for _, sub := range subs {
		go func(s Subscription) {
			for {
				var item Item

				select {
				case item = <-s.Updates():
				case <-m.quit:
					m.err <- s.Close()
					return
				}

				select {
				case m.updates <- item:
				case <-m.quit:
					m.err <- s.Close()
					return
				}
			}
		}(sub)
	}

	return m
}
