package main

import (
	"fmt"
	"strings"
)

var s, t string

func main() {
	fmt.Println("Enter a string:")
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)
	fmt.Println("", strings.Contains(s, "ab"))
	fmt.Println("", strings.ContainsAny(s, "ef"))
	fmt.Println("", strings.Count(s, "a"))
	fmt.Println("", strings.EqualFold(s, t))
}
