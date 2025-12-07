package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func event() {
	fmt.Println("event begin")
	time.Sleep(2 * time.Second)
	fmt.Println("event end")
	close(done)
}

func main() {
	go event()

	select {
	case <- done:
		fmt.Println("received done signal")
	case <- time.After(1 * time.Second):
		fmt.Println("timeout occurred")
		return
	}
}

