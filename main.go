package main

import (
	"fmt"
	"runtime/debug"
)

func read() {
	var list []int = []int {1, 2}
	fmt.Println(list[2])
}

func main1() {
	defer func() {
		fmt.Println("Recovering from panic...")
		err := recover()
		if err != nil {
			fmt.Println("Recovered error:", err)
			fmt.Println("Stack trace:", string(debug.Stack()))
		}
	}()
	read()
}

func main() {
	main1()

	fmt.Println("End of main")
}

