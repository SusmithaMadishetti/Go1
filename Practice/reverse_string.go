package main

import "fmt"

var (
	s string
)

func main() {
	fmt.Print("Enter any string to reverse:")
	fmt.Scan(&s)
	fmt.Print(rev(s))
}
func rev(s string) string {
	slice := []rune(s) //rune acccepts non asci characters too.
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}
