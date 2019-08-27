package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}
type Ver struct {
	a, b float64
}

func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (v *vertex) scale(f float64) { //pointer as receivers-->takes either value or pointer
	v.x = v.x + f
	v.y = v.y + f
}
func abc(x Ver) float64 {
	return math.Sqrt(x.a*x.a + x.b*x.b)
}
func scal(x *Ver, f float64) { //pointer as functions-->takes only &x of ver
	x.a = x.a + f
	x.b = x.b + f
}
func main() {
	v := vertex{3, 4}
	v.scale(5)
	fmt.Println(v.abs())
	x := Ver{4, 3}
	scal(&x, 10)
	fmt.Println(abc(x))
}
