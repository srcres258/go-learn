package main

import "fmt"
import "time"
import "sync"

var moneyChan = make(chan int)

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Println(name, "started shopping")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "finished shopping")

	moneyChan <- money

	wait.Done()
}

func main() {
	start := time.Now()

	var wait sync.WaitGroup

	wait.Add(3)

	go pay("Alpha", 2, &wait)
	go pay("Bravo", 3, &wait)
	go pay("Charlie", 5, &wait)

	go func() {
		wait.Wait()
		close(moneyChan)
	}()

	// for {
	// 	money, ok := <- moneyChan
	// 	fmt.Println(money, ok)
	// 	if !ok {
	// 		break
	// 	}
	// }
	var moneyList []int
	for money := range moneyChan {
		fmt.Println("Received money:", money)
		moneyList = append(moneyList, money)
	}

	fmt.Println("All shopping done, time:", time.Since(start))
	fmt.Println("Money received:", moneyList)
}

