package main

import "fmt"

const pi float64 = 3.14

var r int

func main() {
	fmt.Printf("enter radius of the circle:")
	fmt.Scanf("%d", &r)
	fmt.Println("area: ", pi*float64(r*r)) //type int conversion to float64
	fmt.Println("circum: ", pi*float64(2*r))
}
