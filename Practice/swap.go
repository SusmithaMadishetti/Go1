package main

import "fmt"

//a function can return any number of results.by giving (string,string)
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
