package channels_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/MohammadBnei/go-exercice/channels"
)

func TestBroadcaster(t *testing.T) {
	bc := channels.NewBroadcaster(10)
	defer bc.Close()

	messages := make(chan interface{})
	messages2 := make(chan interface{})

	bc.Register(messages)
	bc.Register(messages2)

	var wg sync.WaitGroup

	wg.Add(4)
	go func(msg <-chan interface{}) {
		i := 1
		for m := range msg {

			sm, ok := m.(string)
			if !ok {
				t.Errorf("wrong type passed to message channel, expected %T got %T", "", sm)
				return
			}
			if sm != fmt.Sprintf("test%v", i) {
				t.Errorf("wrong messages, expected %s got %s", fmt.Sprintf("test%v", i), sm)
			}
			i++
			wg.Done()
		}
		if i != 3 {
			t.Errorf("expected %v run, got %v", 3, i)
		}
	}(messages)

	go func(msg <-chan interface{}) {
		i := 1
		for m := range msg {

			sm, ok := m.(string)
			if !ok {
				t.Errorf("wrong type passed to message channel, expected %T got %T", "", sm)
				return
			}
			if sm != fmt.Sprintf("test%v", i) {
				t.Errorf("wrong messages, expected %s got %s", fmt.Sprintf("test%v", i), sm)
			}
			i++
			wg.Done()
		}
		if i != 4 {
			t.Errorf("expected %v run, got %v", 4, i)
		}
	}(messages2)

	bc.Submit("test1")
	bc.Submit("test2")
	wg.Wait()

	wg.Add(1)
	bc.Unregister(messages)
	bc.Submit("test3")
	wg.Wait()

}

func TestClosedBroadcaster(t *testing.T) {
	bc := channels.NewBroadcaster(10)

	messages := make(chan interface{})

	bc.Register(messages)
	bc.Close()

	if !bc.Submit("test closed") {
		t.Error("should return false, returned true")
	}
}

func TestMultipleBroadcaster(t *testing.T) {
	bc1 := channels.NewBroadcaster(10)

	messages1 := make(chan interface{})

	bc1.Register(messages1)
	defer bc1.Close()

	bc2 := channels.NewBroadcaster(10)

	messages2 := make(chan interface{})

	bc2.Register(messages2)
	defer bc2.Close()

	var wg sync.WaitGroup

	wg.Add(2)

	go func(msg <-chan interface{}) {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			m, ok := <-msg
			if !ok {
				break
			}
			sm, ok := m.(string)
			if !ok {
				t.Errorf("wrong type passed to message channel, expected %T got %T", "", sm)
				return
			}
			if sm != fmt.Sprintf("bc1 %v", i+1) {
				t.Errorf("wrong messages, expected %s got %s", fmt.Sprintf("bc1 %v", i+1), sm)

			}
		}
	}(messages1)

	go func(msg <-chan interface{}) {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			m, ok := <-msg
			if !ok {
				break
			}
			sm, ok := m.(string)
			if !ok {
				t.Errorf("wrong type passed to message channel, expected %T got %T", "", sm)
				return
			}
			if sm != fmt.Sprintf("bc2 %v", i+1) {
				t.Errorf("wrong messages, expected %s got %s", fmt.Sprintf("bc2 %v", i+1), sm)

			}

		}
	}(messages2)

	bc1.Submit("bc1 1")
	bc1.Submit("bc1 2")
	bc1.Submit("bc1 3")

	bc2.Submit("bc2 1")
	bc2.Submit("bc2 2")
	bc2.Submit("bc2 3")

	wg.Wait()
}

func BenchmarkBroadcaster(b *testing.B) {
	bc := channels.NewBroadcaster(1000)
	defer bc.Close()
	done := make(chan bool)

	msgNumber := b.N

	listenerRoutine := func(msg <-chan interface{}) {
		i := 0
		for m := range msg {

			sm, ok := m.(string)
			if !ok {
				b.Errorf("wrong type passed to message channel, expected %T got %T", "", sm)
				return
			}
			if sm != fmt.Sprintf("test%v", i) {
				b.Errorf("wrong messages, expected %s got %s", fmt.Sprintf("test%v", i), sm)
			}
			i++
			done <- true
		}
		if i != msgNumber {
			b.Errorf("expected %v run, got %v", msgNumber, i)
		}
	}

	for i := 0; i < b.N; i++ {
		messages := make(chan interface{})
		bc.Register(messages)
		go listenerRoutine(messages)
	}

	for i := 0; i < b.N; i++ {
		bc.Submit(fmt.Sprintf("test%v", i))
	}
	for i := 0; i < b.N; i++ {
		<-done
	}
}
