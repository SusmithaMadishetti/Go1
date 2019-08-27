package main

import "fmt"

func main() {
	var num = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range num {
		fmt.Println("2**%d=%d\n", i, v)
	}
	for _, v := range num {
		fmt.Println(v)
	}
	for i, _ := range num {
		fmt.Println(i)
	}
	for i := range num {
		fmt.Println(i)
	}
}
