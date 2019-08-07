package main

import (
	"fmt"
	"math"
)

var (
	n             [10]float64
	sum, mean, sd float64
)

func main() {
	fmt.Print("enter 10 numbers:")
	for i := 1; i <= 10; i++ {
		fmt.Scanf("%g", &n[i-1])
		sum += n[i-1]
	}
	mean = sum / 10
	for i := 0; i < 10; i++ {
		sd += math.Pow(n[i]-mean, 2)
	}
	fmt.Print("sd:", math.Sqrt(sd/10))
}
