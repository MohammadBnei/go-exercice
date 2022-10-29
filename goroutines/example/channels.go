package goroutines

import "fmt"

var done chan bool = make(chan bool)

func f5() {
	fmt.Println("func 5")
	done <- true
}
func f6() {
	fmt.Println("func 6")
	done <- true
}

func CRun() {
	fmt.Println("Wait Group Start...")
	go f5()
	go f6()
	<-done
	<-done
	fmt.Println("End")
}
