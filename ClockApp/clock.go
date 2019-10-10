package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	message, y, z string
	m             map[int]string
)

func main() {
	TickTockBong()
}
func updateMessage(s *string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		*s = scanner.Text()
	}
}

//TickTockBong returns map
func TickTockBong() map[int]string {
	i := 0
	message = "tick"
	y = "tock"
	z = "bong"
	m = make(map[int]string)
	go updateMessage(&message)
	for {
		i++
		time.Sleep(1 * time.Second)
		switch {
		case i%(3*60*60) == 0: //3 hour case
			fmt.Println(z, i)
			m[i] = z
			return m
		case i%(int(60*60)) == 0: //every hour case
			fmt.Println(z, i)
			m[i] = z
		case i%(60) == 0: //minute case
			fmt.Println(y, i)
			m[i] = y
		default:
			fmt.Println(message, i)
			m[i] = message
		}
	}
}
