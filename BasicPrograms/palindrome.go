package main

import "fmt"

var (
	n, rem, rev int
)

func main() {
	fmt.Printf("Enter a number to check if it palindrome:")
	fmt.Scanf("%d", &n)
	temp := n
	for n > 0 {
		rem = n % 10
		rev = rev*10 + rem
		n /= 10
	}
	if temp == rev {
		fmt.Printf("%d is palindrome", temp)
	} else {
		fmt.Printf("%d is not a palindrome", temp)
	}
}
