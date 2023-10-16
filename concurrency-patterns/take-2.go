package main

import (
	"fmt"
	"sync"
)

func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- id*10 + i
	}
}
func take(ch <-chan int, num int) []int {
	var items []int
	for i := 0; i < num; i++ {
		select {
		case item, ok := <-ch:
			if !ok {
				return items
			}
			items = append(items, item)
		}
	}
	return items
}

func main() {
	ch := make(chan int, 50)
	var wg sync.WaitGroup
	numProducers := 5
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go producer(i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := 0; i < numProducers; i++ {
		items := take(ch, 3)
		fmt.Println(items)
	}
}
