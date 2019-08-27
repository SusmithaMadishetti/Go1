package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	i := strings.Fields(s)
	m := make(map[string]int)
	fmt.Println(i)
	for _, words := range i {
		m[words]++
	}
	return m

}

func main() {
	wc.Test(WordCount)
}
