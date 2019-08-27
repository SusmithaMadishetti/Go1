package main

import "fmt"

// type stringer interface {
// 	so() string
// }
type user struct {
	name string
	age  int
}

func (p user) string() string {
	return fmt.Sprintf("%v(%v yeras)", p.name, p.age)
}
func main() {
	a := user{"pooja", 12}
	fmt.Println(a.string())
}
