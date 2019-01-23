// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		case update := <- add_number:
			i += update

		case signal := <- control:
			if signal == GetNumber{
				number <- i
			}
			if signal == Exit{
				return
			}
			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.
		}
	}
}

func incrementing(add_number chan<-int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- 1
	}
	finished <- true
	//TODO: signal that the goroutine is finished
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- -1
	}
	finished <- true
	//TODO: signal that the goroutine is finished
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	finished := make(chan bool, 2)
	number := make(chan int)
	add_number:= make(chan int)
	control := make(chan int)
	// TODO: Construct the required channels
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.

	go number_server(add_number, control, number)
	go incrementing(add_number, finished)
	go decrementing(add_number, finished)
	// TODO: Spawn the required goroutines

	<- finished
	<- finished
	// TODO: block on finished from both "worker" goroutines

	control<-GetNumber
	Println("The magic number is:", <- number)
	control<-Exit
}
