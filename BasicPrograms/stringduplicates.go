package main

import (
	"fmt"
	"strings"
)

var result []string

func main() {
	st := "application requires knowledge"
	fmt.Println(st)
	words := strings.Fields(st)
	for i, v := range words {
		fmt.Println(i, v)
		char := []rune(v)
		r := string(sorted(char))
		fmt.Println(r)
		result = append(result, r)
	}
	fmt.Println(result)
	fmt.Println(strings.Join(result, " "))
}
func sorted(char []rune) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, e := range char {
		if _, v := keys[e]; !v {
			keys[e] = true
			list = append(list, e)
		}
	}
	return list
}
