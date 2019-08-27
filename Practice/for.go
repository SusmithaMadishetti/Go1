package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//here initialization and incrementation is optional this is nothing but while in GO
	// for ;condtion;{}
	j := 2
	for j < 5 {
		j += j
	}
	fmt.Print(j)
}
