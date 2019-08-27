package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//T(v) converts value v to type T
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//we can write in this way without declaring value
	var z = uint(f)
	fmt.Println(x, y, z)
	v := 42.5 + 65i // this declaration of variable takes the type from right hand side
	fmt.Printf("v is of type %T\n", v)
}
