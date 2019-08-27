package main

import "fmt"

type user struct {
	name, anumber string
}

var m, k map[string]user

func main() {
	m = make(map[string]user)
	m["1"] = user{"sush", "321"}
	m["2"] = user{"karth", "4556"}
	fmt.Println(m)
	fmt.Println(m["2"])
	n := map[string]user{
		"1": user{"abc", "1234"},
		"2": {"def", "5678"}, // we can define without giving type user
	}
	fmt.Println(n)
	m["3"] = user{"ghi", "1111"}
	fmt.Println("value:", m["3"])
	delete(k, "1")
	fmt.Println(k)
	v, ok := k["0"] // v is nothin but saving into the variable v.of key 2
	fmt.Println("The value:", v, "Present?", ok)
}
