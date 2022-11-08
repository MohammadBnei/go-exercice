package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/MohammadBnei/go-exercice/channels"
)

func main() {
	in := make(chan string)
	out := make(chan string)
	channels.NewSimpleMessaging(in, out)

	var wg sync.WaitGroup
	wg.Add(20)
	go func() {
		for i := 0; i < 10; i++ {
			in <- fmt.Sprintf("message #%v", i)
			time.Sleep(time.Second)
			wg.Done()
		}
	}()

	go func(out <-chan string) {
		for r := range out {
			fmt.Println(r)
			wg.Done()
		}
	}(out)
	wg.Wait()
	// close(out)
	// close(in)
}
