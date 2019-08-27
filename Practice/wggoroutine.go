package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func a() {
	fmt.Println("a")
	wg.Done()

}
func main() {
	wg.Add(1)
	go a()
	wg.Wait()
	fmt.Println("main")
}
