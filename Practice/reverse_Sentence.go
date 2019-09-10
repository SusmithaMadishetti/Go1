package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var s string

func main() {
	fmt.Println("Enter the sentence to reverse:")
	consoleReader := bufio.NewReader(os.Stdin)
	s, _ := consoleReader.ReadString('\n')
	//fmt.Scanf(&s) // Scanf,Fscanf,Sscanf requires newlines in the input to match newlines in format. Scanln,Fscanln,Sscanln stop scanning at a newline and require that the items to be followed by one
	fmt.Println(reve(s))
}
func reve(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
