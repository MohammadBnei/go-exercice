package goroutines

import (
	"fmt"
)

func f1() {
	fmt.Println("func 1")
}
func f2() {
	fmt.Println("func 2")
}

func FJRun() {
	fmt.Println("Fork-Join Start...")
	go f1()
	go f2()
	fmt.Println("End")
}
