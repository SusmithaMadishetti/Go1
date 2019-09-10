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
func b() {
	fmt.Println("b")
	wg.Done()

}
func main() {
	wg.Add(2)
	fmt.Println("before  b")

	go b()
	fmt.Println("after b")

	fmt.Println("after bbb")
	go a()
	//fmt.Println("after a and before wait")

	wg.Wait()
	fmt.Println("main")
}
