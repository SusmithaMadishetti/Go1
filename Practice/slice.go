package main

import "fmt"

var (
	num [10]int
)

func main() {
	fmt.Println("enter numbers in array:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&num[i])
	}
	fmt.Println(num)
	var s []int = num[1:10]
	fmt.Println(s, len(s), cap(s))
	s[1] = 44 //they dont store,they can be modified.
	fmt.Println(s, len(s), cap(s))
	fmt.Println(num)
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	a = a[2:6]
	fmt.Println(a, len(a), cap(a))
	a = a[:2]
	fmt.Println(a, len(a), cap(a))
	a = a[1:]
	fmt.Println(a, len(a), cap(a))
	var b []int
	fmt.Println(b, len(b), cap(b))
	if b == nil {
		fmt.Println("nil!")
	}
	b = append(b, 0, 1, 2)
	fmt.Println("hey", b, len(b), cap(b))
	//bd := delete(b, 1) we cannot delete in slice
	//fmt.Println(b)
	c := make([]int, 5, 5)
	fmt.Println(c, len(c), cap(c))
	d := c[:2]
	fmt.Println(d, len(d), cap(d))
}
