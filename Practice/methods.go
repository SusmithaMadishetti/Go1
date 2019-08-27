package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x float64
	y float64
}
type myint int // other than struct

func (v vertex) sqrt() float64 { //methods like class using receiver argument
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func sqrt(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func (i myint) ab() int {
	if i < 0 {
		return int(-i)
	}
	return int(i)
}
func main() {
	v := vertex{3, 4}
	i := myint(math.Pow(-2, 2)) //methods as class
	fmt.Println(v.sqrt())
	fmt.Println(sqrt(v)) //methods as function
	fmt.Println(i.ab())
}
