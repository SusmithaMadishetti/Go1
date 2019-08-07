package main

import "fmt"

var n, i int

func main() {
	fmt.Println("enter the number for multiplication:")
	fmt.Scanf("%d", &n)
	i := 0
	for i <= 10 {
		fmt.Printf("%d * %d = %d", n, i, n*i)
		fmt.Println(" ")
		i++
	}
}
