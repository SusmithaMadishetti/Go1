package main

import (
	"fmt"
	"math/cmplx"
)

//variables of several types,and also that variable declarations may be "factored" into blocks,
// as with import statements.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	sus    string     = "susmitha"
	//default values for various data types
	i int
	f float64
	b bool
	s string
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println("Type: %q Value: %q\n", sus, sus)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
