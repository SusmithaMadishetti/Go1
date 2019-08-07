package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter a number to check even/odd:")
	fmt.Scanf("%d", &n)
	if n%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
