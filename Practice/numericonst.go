package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//first takes 2 and returns 21
	fmt.Println(needInt(Small))
	//second takes 2 and returns 0.2
	fmt.Println(needFloat(Small))
	//third takes 0.123456.... 100 places and returns 1.23456...
	fmt.Println(needFloat(Big))
}
