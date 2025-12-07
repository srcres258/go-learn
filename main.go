package main

import "fmt"
import "time"
import "sync"

func shopping(name string, wait *sync.WaitGroup) {
  fmt.Println(name, "started shopping")
  time.Sleep(1 * time.Second)
  fmt.Println(name, "finished shopping")
  wait.Done()
}

func main() {
  start := time.Now()
	var wait sync.WaitGroup
  wait.Add(4)
  go shopping("Alpha", &wait)
  go shopping("Bravo", &wait)
  go shopping("Charlie", &wait)
  go shopping("Delta", &wait)
  wait.Wait()
  fmt.Println("All shopping done, time:", time.Since(start))
}

