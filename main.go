package main

import "fmt"

var maps = map[int]string{}

func main() {
	go func() {
		for {
			maps[1] = "foo"
		}
	}()
	go func() {
		for {
			fmt.Println(maps[1])
		}
	}()

	select{}
}

