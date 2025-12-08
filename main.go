package main

import (
	"fmt"
	"os"
)

func init() {
	_, err := os.ReadFile("111")
	if err != nil {
		panic("错误了")
	}
}

func main() {
	fmt.Println("Hello, World!")
}

