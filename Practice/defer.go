package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("deferred")
	defer fmt.Println("zello")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}

	fmt.Println("hello")
}
