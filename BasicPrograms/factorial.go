package main

import "fmt"

var (
	i, res, n int
)

func fact(n int) int {
	res := 1
	for i = 1; i <= n; i++ {
		res = res * i
	}
	return res
}
func main() {
	fmt.Print("Enter number:")
	fmt.Scanf("%d", &n)
	fmt.Println("factorial:", fact(n))
}
