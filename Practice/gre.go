package main

import "fmt"

func greater(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	fmt.Println(greater(10, 3), greater(2, 7))
}
