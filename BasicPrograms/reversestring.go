package main

import "fmt"

var str string

func main() {
	a := []byte{}
	fmt.Println("enter string:")
	fmt.Scanf("%s", &str)
	b := []byte(str)
	for i := 1; i < len(b)+1; i++ {
		a = append(a, b[len(b)-i])
	}
	fmt.Print(a)
	fmt.Println(string(a))
}
