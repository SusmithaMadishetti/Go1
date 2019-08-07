package main

import (
	"fmt"
	"strings"
)

var (
	a, b, c, p string
)

func main() {
	fmt.Println("Enter two strings:")
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	p := "hello"
	s := []string{"sus", "kar", "abc"}
	fmt.Println("", strings.Contains(a, "bc"))
	fmt.Println("", strings.Count(a, "b"))
	fmt.Println("", strings.HasPrefix(a, "ab"))
	fmt.Println("", strings.HasSuffix(a, "c"))
	fmt.Println("", strings.Index(a, "o"))
	fmt.Println("", strings.Join(s, p))
	fmt.Println("", strings.Repeat(b, 2))
	fmt.Println("", strings.Replace(a, b, c, 2))
	fmt.Println("", strings.Split(a, b))
	fmt.Println("", strings.ToLower(a))
	fmt.Println("", strings.ToUpper(a))
}
