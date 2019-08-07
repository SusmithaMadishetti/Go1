package main

import (
	"fmt"
	"sort"
)

var (
	a []int
	b []string
	c []float64
)

func main() {
	fmt.Println("Enter integers to be sorted:")
	for i := 0; i < 5; i++ {

		fmt.Scan(a[i])
	}
	fmt.Println("", sort.IntsAreSorted(a))
}
