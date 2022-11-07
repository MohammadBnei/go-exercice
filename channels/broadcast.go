package channels

type Broadcaster interface {
	// Register a new channel to receive broadcasts
	Register(chan<- interface{})
	// Unregister a channel so that it no longer receives broadcasts.
	Unregister(chan<- interface{})
	// Shut this broadcaster down.
	Close() error
	// Submit a new object to all subscribers
	Submit(interface{}) bool
}

type broadcaster struct {
	input chan interface{}
	reg   chan chan<- interface{}
	unreg chan chan<- interface{}

	outputs map[chan<- interface{}]bool
}

func NewBroadcaster(buflen int) Broadcaster {
	return &broadcaster{}
}

func (b *broadcaster) Register(chan<- interface{})   {}
func (b *broadcaster) Unregister(chan<- interface{}) {}
func (b *broadcaster) Close() error {
	return nil
}
func (b *broadcaster) Submit(interface{}) bool {
	return true
}
