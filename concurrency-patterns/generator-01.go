package main

import "log"

func main() {
	generator := func(done <-chan interface{}, num ...int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for _, i := range num {
				select {
				case <-done:
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}
	done := make(chan interface{})
	defer close(done)
	collection := generator(done, 1, 2, 3, 4, 5, 6)
	for i := range collection {
		log.Println(i)
	}
}
