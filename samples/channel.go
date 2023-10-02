package main

import "sync"

type Counter struct {
	count int
	mutex chan struct{}
}

func NewCounter() *Counter {
	return &Counter{
		mutex: make(chan struct{}, 1),
	}
}
func (c *Counter) Inc() {
	c.mutex <- struct{}{}
	c.count++
	<-c.mutex
}
func (c *Counter) Count() int {
	c.mutex <- struct{}{}
	defer func() { <-c.mutex }()
	return c.count
}
func main() {
	var wg sync.WaitGroup
	counter := NewCounter()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	println("Counter:", counter.Count())
}
