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
	return nil
}

func (m *merge) Updates() <-chan Item {
	return m.updates
}

// Merge follows the Fan-In pattern and merges
// multiple Subscriptions into one
func Merge(subs ...Subscription) Subscription {
	return nil
}
