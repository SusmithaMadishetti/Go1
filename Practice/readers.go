package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("susmithamadishetti")
	b := make([]byte, 4)
	for {
		n, err := r.Read(b)
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}