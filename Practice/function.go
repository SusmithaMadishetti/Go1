package main

import (
	"fmt"
	"math"
)

func compute(f func(float64, float64) float64) float64 {
	return f(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	sqrt := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(sqrt(5, 12))
	fmt.Println(compute(sqrt)) // function as value
	fmt.Println(compute(math.Pow))
	pos := adder() // function closure can access and assign values of referenced variables
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i))
	}
}
