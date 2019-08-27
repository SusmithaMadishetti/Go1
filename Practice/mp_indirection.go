package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v *vertex) scale(f float64) { //pointer should take pointer as receiver
	v.x = v.x * f
	v.y = v.y * f
}
func scalefunc(v *vertex, f float64) { //pointer should take pointer as argument
	v.x = v.x * f
	v.y = v.y * f
}
func (v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func absfunc(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main() {
	v := vertex{3, 4}
	v.scale(2)         //reciever
	scalefunc(&v, 10)  //pointer for v
	p := &vertex{4, 3} // pointer for p (it is better to use pointers)
	p.scale(3)
	scalefunc(p, 8)
	fmt.Println(v, p)
	//v := vertex{3, 4}
	fmt.Println(v.abs()) //methods as value
	fmt.Println(absfunc(v))
	//p := &vertex{4, 3}
	fmt.Println(p.abs()) //method as pointer
	fmt.Println(absfunc(*p))
}
