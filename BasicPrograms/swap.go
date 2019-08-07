package main

import (
	"fmt"
)

var (
	x, y int
)

func swap(x, y int) (int, int) {
	return y, x
}
func main() {
	fmt.Printf("Enter 2 numbers to swap:")
	fmt.Scanf("%d%d", &x, &y)
	i, j := swap(x, y)
	fmt.Printf("numbers after swapping:%d %d", i, j)
}
