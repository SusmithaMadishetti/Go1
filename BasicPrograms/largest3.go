package main

import "fmt"

func largest(x, y, z int) int {
	if x > y && x > z {
		return x
	}
	if y > x && y > z {
		return y
	}
	return z
}
func main() {
	fmt.Println("enter 3 number:")
	var a, b, c int
	fmt.Scanf("%d%d%d", &a, &b, &c)
	fmt.Println("largest", largest(a, b, c))
}
