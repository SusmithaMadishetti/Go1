package main

import (
	//multiple packages can be written in brackets
	"fmt"
	"math"
	"math/rand"
	//sub packages can be written in main/subpackage
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Square root of 7 is", math.Sqrt(7))
	//always package.Captialiaze(function)
}
