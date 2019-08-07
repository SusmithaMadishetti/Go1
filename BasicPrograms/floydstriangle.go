package main

import "fmt"

var (
	n, i, j, k int
)

func main() {
	fmt.Print("enter number of rows in floyd's traingle:")
	fmt.Scanf("%d", &n)
	k = 1
	for i = 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Printf("%d", k)
			k++
		}
		fmt.Println("")
	}
}
