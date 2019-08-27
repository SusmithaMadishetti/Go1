package main

import "fmt"

func main() {
	var i, j = 42, 100
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p, &p, &i, i)
	p = &j
	*p = *p / 50
	fmt.Println(j)
}
