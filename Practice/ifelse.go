package main

import (
	"fmt"
	"math"
)

func power(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func main() {
	fmt.Println(
		power(4, 2, 10),
		power(2, 2, 10),
	)
}
