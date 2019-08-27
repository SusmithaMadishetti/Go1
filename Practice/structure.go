package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	fmt.Println(student{4, "sush"})
	s := student{5, "karth"}
	s1 := student{id: 7}
	s2 := student{}
	p1 := &student{4, "sush"}
	fmt.Println(s.id, s.name)
	p := &s
	p.id = 34
	fmt.Println(s, s1, s2, p1)
}
