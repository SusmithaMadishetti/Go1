package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}
type Myfloat float64
type abser interface { //creating an interface should implement all the abstract methods
	abs() float64
	//abc() int
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (f Myfloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	var a abser
	f := Myfloat(-math.Sqrt(2))
	v := vertex{3, 4}
	a = f // implementing abstract methods
	fmt.Println(a.abs())
	a = &v //implemneting abstract methods
	fmt.Println(a.abs())
}
