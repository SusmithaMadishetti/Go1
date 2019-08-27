package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func print(p int) {
	for i := 6; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)

		fmt.Println(p)
	}
}
func main() {
	go say("hello")
	go say("world")
	go print(5)
	say("function Go Lang")
	fmt.Println("im in print")
}
