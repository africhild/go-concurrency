## Lesson 2: Go Concurrency Patterns

### [Take](take.go): 
The concept of "take" is about retrieving a certain number of items from a channel until that number is met or the channel is closed. In Go, you'd typically use a for loop and a select statement to achieve this.
 ```go
package main

import (
	"fmt"
)

func take(ch <-chan int, num int) []int {
	var items []int

	for i := 0; i < num; i++ {
		select {
		case item, ok := <-ch:
			if !ok {
				return items // Return early if the channel is closed
			}
			items = append(items, item)
		}
	}

	return items
}

func main() {
	// An example channel with some items
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // It's a good practice to close channels when you're done sending

	// Using the take function to get the first 3 items
	items := take(ch, 3)
	fmt.Println(items) // This will print: [0 1 2]
}
```

```go
package main

import (
	"fmt"
	"sync"
)

// Producer: produces items and sends them to the channel.
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- id*10 + i
	}
}

// Take: retrieves a subset of items from the channel.
func take(ch <-chan int, num int) []int {
	var items []int

	for i := 0; i < num; i++ {
		select {
		case item, ok := <-ch:
			if !ok {
				return items // Return early if the channel is closed
			}
			items = append(items, item)
		}
	}

	return items
}

func main() {
	ch := make(chan int, 50) // buffer size to hold all the items

	var wg sync.WaitGroup
	numProducers := 5
	for i := 1; i <= numProducers; i++ {
		wg.Add(1)
		go producer(i, ch, &wg)
	}

	// Close the channel after all producers are done sending data
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Using the take function multiple times to get subsets of items
	for i := 0; i < numProducers; i++ {
		items := take(ch, 3) // Take 3 items at a time
		fmt.Println(items)
	}
}

```