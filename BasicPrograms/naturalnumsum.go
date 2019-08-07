package main

import (
	"fmt"
)

var n, i int

func main() {
	fmt.Printf("Enter the number for natural sum:")
	fmt.Scanf("%d", &n)
	sum := 0
	for i = 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("sum :%d", sum)
}
