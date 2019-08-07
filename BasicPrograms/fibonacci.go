package main

import "fmt"

var (
	n, n3 int
)

func main() {
	n1 := 0
	n2 := 1
	fmt.Println("enter a number to generate fibonacci:")
	fmt.Scan(&n)
	fmt.Print(n1)
	fmt.Print(n2)
	for i := 0; i < n-2; i++ {
		n3 = n1 + n2
		n1 = n2
		n2 = n3
		fmt.Print(n3)
	}

}
