package main

import (
	"fmt"
	"sync"
)

func main() {
	var maps = sync.Map{}
	go func() {
		for {
			maps.Store(1, "hello")
		}
	}()
	go func() {
		for {
			v, ok := maps.Load(1)
			fmt.Println(v, ok)
		}
	}()

	select{}
}

