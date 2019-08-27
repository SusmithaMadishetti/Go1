package main

import "fmt"
//constants are not declared by using :=
//they can be any value string,boolean,numeric or character
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}