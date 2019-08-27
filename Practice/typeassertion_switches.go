package main

import "fmt"

func do(a interface{}) { //type switches
	switch v := a.(type) {
	case int:
		fmt.Printf("%v %T", v, v)
	case string:
		fmt.Printf("%q %T", v, v)
	default:
		fmt.Printf("%T", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(3.14)
	var i interface{} = "hello"
	var j interface{} = 2
	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)
	//f := i.(float64) // panic occured as i doesnt hold any type of float64
	//fmt.Println(f)
	//f, ok = i.(float64)
	//fmt.Println(f, ok)
	k := j.(int)
	fmt.Println(k)
	k, ok = j.(int)
	fmt.Println(k, ok)

}
