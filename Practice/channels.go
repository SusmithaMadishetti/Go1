package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 0, -8, -7}
	ch := make(chan int) //create chan before initializing
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)
	x := <-ch
	y := <-ch
	fmt.Println(x, y, x+y)
	chbuf := make(chan int, 3) // you can overfil the chanels
	chbuf <- 1
	chbuf <- 2
	chbuf <- 3
	fmt.Println(ch, ch, ch)
}

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}
