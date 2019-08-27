package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (a *rot13Reader) Read(s []byte) (n int, err error) {
	n, err = a.r.Read(s)
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'A' && s[i] < 'N') || (s[i] >= 'a' && s[i] < 'n') {
			s[i] += 13
		} else if (s[i] > 'M' && s[i] <= 'Z') || (s[i] > 'm' && s[i] <= 'z') {
			s[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("fhfzvgun")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
