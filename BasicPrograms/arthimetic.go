package main

import "fmt"

var (
	a, b int
)

func main() {
	fmt.Printf("Enter two number:")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println("add:", a+b)
	fmt.Println("sub:", a-b)
	fmt.Println("mult:", a*b)
	fmt.Println("div:", a/b)
	fmt.Println("rem:", a%b)
}
