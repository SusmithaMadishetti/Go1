package main

import "fmt"

//default boolean value is false
var c, python, java bool

//If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var cp, python2, java8 = true, false, "no!"

func main() {
	//deafult int value is 0
	var i int
	a := 3 //outside function we cant declare short variables
	var k, j int = 1, 2
	fmt.Println(i, a, j, k, c, python, java)
	fmt.Println(cp, python2, java8)
}
