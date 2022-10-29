package goroutines

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f3() {
	defer wg.Done()
	fmt.Println("func 3")
}
func f4() {
	defer wg.Done()
	fmt.Println("func 4")
}

func WGRun() {
	fmt.Println("Wait Group Start...")
	wg.Add(2)
	go f3()
	go f4()
	fmt.Println("End")
	wg.Wait()
}
