package channels

type Broadcaster[T any] interface {
	// Register a new channel to receive broadcasts
	Register(chan<- T)
	// Unregister a channel so that it no longer receives broadcasts.
	Unregister(chan<- T)
	// Shut this broadcaster down.
	Close() error
	// Submit a new object to all subscribers
	Submit(T) bool
}

type broadcaster[T any] struct {
	input chan T
	reg   chan chan<- T
	unreg chan chan<- T

	outputs map[chan<- T]bool
}

func NewBroadcaster[T any](buflen int) Broadcaster[T] {
	return &broadcaster[T]{}
}

func (b *broadcaster[T]) Register(chan<- T)   {}
func (b *broadcaster[T]) Unregister(chan<- T) {}
func (b *broadcaster[T]) Close() error {
	return nil
}
func (b *broadcaster[T]) Submit(T) bool {
	return true
}
