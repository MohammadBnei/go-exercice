package channels

type simpleMessaging struct {
	input  <-chan string
	output chan<- string
}

func (sm *simpleMessaging) run() {
	for m := range sm.input {
		sm.output <- m
	}
}

func NewSimpleMessaging(input, output chan string) *simpleMessaging {
	sm := &simpleMessaging{input, output}
	go sm.run()
	return sm
}

func (sm *simpleMessaging) Submit(m string) {
	sm.output <- m
}
