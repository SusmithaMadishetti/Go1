package main

import (
	"fmt"
	"math"
)

const pi = 3.14

type polygon interface {
	area()
}

// type empty interface{

// }

type circle struct {
	r float64
}
type square struct {
	s float64
}
type rectangle struct {
	l, b float64
}
type fl float64
type T struct {
	s string
}

func (f fl) area() {
	fmt.Println(f)
}

func (t *T) area() {
	if t == nil {
		fmt.Println("nil!")
		return
	}
	fmt.Println(t.s)
}
func describe(p polygon) {
	fmt.Printf("%v,%T", p, p)
}
func describe1(em interface{}) {
	fmt.Printf("%v %T", em, em)
}

func (c circle) area() {
	fmt.Println(2 * pi * c.r)
}
func (side square) area() {
	fmt.Println(side.s * side.s)
}
func (len rectangle) area() {
	fmt.Println(len.l * len.b)
}

func main() {
	i := circle{5}
	i.area()
	j := square{4}
	j.area()
	k := rectangle{3, 4}
	k.area()
	describe(k)
	l := fl(math.Pi) // interface value
	describe(l)
	l.area()
	var n *T // interface value having nil value
	describe(n)
	n.area()

	m := &T{"hello"} // interfcae pointer  question why is it not taking address value
	describe(m)
	m.area()

	var em interface{}
	describe1(em) // empty interface
	em = 42
	describe1(em)
	em = "heya"
	describe1(em)
	g := []polygon{i, j, k}
	for p := range g {
		fmt.Println(g[p].area())
	}
}
