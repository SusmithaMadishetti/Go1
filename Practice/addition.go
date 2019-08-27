package main

import "fmt"
//takes value of int type arguments and returns int type
//if both parameters are sharing the same data type then we can write (x,y int) int
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
