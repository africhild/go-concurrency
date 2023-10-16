package main

import "fmt"

func take(ch <-chan int, item int) []int {
	var items []int
	for i := 0; i < item; i++ {
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
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	item := take(ch, 3)
	fmt.Println(item)
}
