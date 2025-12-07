package main

import "fmt"
import "time"
import "sync"

var moneyChan = make(chan int)
var nameChan = make(chan string)
var doneChan = make(chan struct{})

func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Println(name, "started shopping")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "finished shopping")

	moneyChan <- money
	nameChan <- name

	wait.Done()
}

func main() {
	start := time.Now()

	var wait sync.WaitGroup

	wait.Add(3)

	go send("Alpha", 2, &wait)
	go send("Bravo", 3, &wait)
	go send("Charlie", 5, &wait)

	go func() {
		wait.Wait()
		close(moneyChan)
		close(nameChan)
		close(doneChan)
	}()

	// for {
	// 	money, ok := <- moneyChan
	// 	fmt.Println(money, ok)
	// 	if !ok {
	// 		break
	// 	}
	// }
	var moneyList []int
	var nameList []string
	// go func() {
	// 	for money := range moneyChan {
	// 		fmt.Println("Received money:", money)
	// 		moneyList = append(moneyList, money)
	// 	}
	// }()
	// for name := range nameChan {
	// 	nameList = append(nameList, name)
	// }
	var event = func () {
		for {
			select {
			case money := <- moneyChan:
				moneyList = append(moneyList, money)
			case name := <- nameChan:
				nameList = append(nameList, name)
			case <- doneChan:
				return
			}
		}
	}
	event()
	fmt.Println("All shopping done, time:", time.Since(start))
	fmt.Println("Money received:", moneyList)
	fmt.Println("Names received:", nameList)
}

