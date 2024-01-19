package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup, number int) {
	wg.Add(1)
	fmt.Println("New routine", number)
	wg.Done()
}

func main() {
	// c := make(chan int)
	// c <- 3
	// fmt.Println("channel", c)
	// var x int = <-c
	// fmt.Println("channel after", x)
	// close(c)

	var wg sync.WaitGroup

	go foo(&wg, 1)
	go foo(&wg, 2)
	go foo(&wg, 3)

	go foo(&wg, 4)
	go foo(&wg, 5)

	go foo(&wg, 6)
	wg.Wait()

	fmt.Println("Main routine")

}
