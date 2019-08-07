package main

import "fmt"

var num, rem, temp, sum int

func main() {
	fmt.Printf("enter number to check armstrong:")
	fmt.Scanf("%d", &num)
	temp = num
	sum = 0
	for temp > 0 {
		rem = temp % 10
		sum += (rem * rem * rem)
		temp = temp / 10
	}
	if num == sum {
		fmt.Printf("%d is armstrong number", num)
	} else {
		fmt.Printf("%d is not an armstrong number", num)
	}

}
