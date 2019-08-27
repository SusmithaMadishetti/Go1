package main

import (
	"fmt"
	"sync"
	"time"
)

type Safe struct {
	v   map[string]int
	mux sync.Mutex
}

func main() {
	c := Safe{v: make(map[string]int)}
	go c.Inc("hello")
	go c.Inc("world")
	fmt.Println(c.Value("hello"))
	time.Sleep(time.Second)
	fmt.Println(c.Value("world"))
}
func (c *Safe) Inc(s string) {
	c.mux.Lock()
	c.v[s]++
	c.v[s]++
	c.mux.Unlock()
}
func (c *Safe) Value(s string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[s]
}
