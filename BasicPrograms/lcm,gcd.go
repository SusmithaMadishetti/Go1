package main

import "fmt"

var (
	a, b int
	s    int
)

func options() {
	fmt.Println("Enter the option:")
	fmt.Println("1.LCM")
	fmt.Println("2.GCD")
}
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}
func gcd(a, b int) int {
	var v int
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			v = i
		}
	}
	return v
}
func main() {
	fmt.Println("option:1.LCM and 2.GCD")
	fmt.Println("enter two numbers:")
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &s)
	switch s {
	case 1:
		fmt.Print("LCM:", lcm(a, b))
	case 2:
		fmt.Print("GCD:", gcd(a, b))
	}
}
