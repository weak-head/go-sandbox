package feedreader

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

func (s *sub) loop() {

}
