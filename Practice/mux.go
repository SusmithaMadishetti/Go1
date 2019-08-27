package main

import (
	"fmt"
	"runtime"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	fmt.Println("INC")
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Inc1 increments the counter for the given key.
func (c *SafeCounter) Inc1(key string) {
	fmt.Println("INC1")
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++

}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	fmt.Println("Value")
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	runtime.GOMAXPROCS(4)
	c := SafeCounter{v: make(map[string]int)}
	go c.Inc1("somekey")
	// for i := 0; i < 10; i++ {
	// 	go c.Inc("somekey")
	// }
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	go c.Inc1("somekey")
	//time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
