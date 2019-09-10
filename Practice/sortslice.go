package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 4, 2, 6, 4, 8, 5}
	a := []string{"dab", "a"}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(a)

	sort.Strings(a) // sorts strings incresing order.

	fmt.Println(a)

}
