package main

import "fmt"

var (
	a, b int
)

func main() {
	fmt.Printf("enter two numbrs to add:")
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println("add:", a+b)
}
