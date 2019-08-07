package main

import "fmt"

var (
	area int
)

func main() {
	var l, b int
	fmt.Println("enter length and breadth:")
	fmt.Scanf("%d", &l)
	fmt.Scanf("%d", &b)
	fmt.Println("area of rectangle:", l*b)
	fmt.Println("area of a square of length:", l*l)

}
