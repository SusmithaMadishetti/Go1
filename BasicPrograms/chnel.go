package main

import "fmt"

func main() {
	c := make(chan int, 10)
	for i := 0; i < 11; i++ {
		c <- i
		fmt.Println(<-c)
	}
}
