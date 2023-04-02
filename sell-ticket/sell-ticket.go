package main

import (
	"fmt"
	"sync"
)

func sellTickets(wg *sync.WaitGroup, ch chan int, id int) {
	defer wg.Done()

	for {
		ticket, ok := <-ch
		if !ok {
			fmt.Printf("grountine %d: channel is closed\n", id)
			return
		}
		fmt.Printf("grountine %d: sell ticket %d\n", id, ticket)
	}
}
func main() {
	const numTIckets = 100
	const numSellers = 4
	var wg sync.WaitGroup
	wg.Add(numSellers)

	ch := make(chan int, numTIckets)
	for i := 1; i <= numTIckets; i++ {
		ch <- i
	}

	close(ch)

	for i := 1; i <= numSellers; i++ {
		go sellTickets(&wg, ch, i)
	}

	wg.Wait()
}
