package main

import "fmt"

var (
	a   [5]int
	b   [5]int
	sum [5]int
	sub [5]int
	i   int
)

func main() {
	fmt.Println("enter numbers in array:")
	//var n int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println("enter second array:")
	for i = 0; i < 5; i++ {
		fmt.Scan(&b[i])
	}
	fmt.Println(a)
	fmt.Println(b)
	for i = 0; i < 5; i++ {
		sum[i] = a[i] + b[i]
	}
	fmt.Println(sum)
	for i = 0; i < 5; i++ {
		sub[i] = a[i] - b[i]

	}
	fmt.Println(sub)

}
