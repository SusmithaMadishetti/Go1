package main

import (
	"fmt"

	"./add"
	"./div"
	"./mul"
)

func main() {
	var a, b int
	fmt.Println("im in calcu")
	fmt.Println("enter two numbers:")
	fmt.Scanf("%d\n%d", &a, &b)
	fmt.Println(add.A1(a, b))
	fmt.Println(mul.M1(a, b))
	fmt.Println(div.D1(a, b))
}
