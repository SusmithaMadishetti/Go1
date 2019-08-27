package main

import "fmt"

func main() {

	for n := 1; n < 3; n++ {
		fmt.Println("enter fav number")
		fmt.Scanf("%d", &n)
		switch n {
		case 1:
			fmt.Println("appple")
			fallthrough
		case 2:
			fmt.Println("banana")
			fallthrough
		default:
			fmt.Println("default")

		}
	}
}
