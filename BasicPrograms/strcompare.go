package main

import (
	"fmt"
	"strings"
)

var (
	a, b string
)

func main() {
	fmt.Println("enter two strings:")
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	fmt.Println("", strings.Compare(a, b))
}
