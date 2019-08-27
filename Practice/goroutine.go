package main

import (
	"fmt"
	"runtime"
)

func leraning() {
	fmt.Println("learing")
}

func main() {
	runtime.GOMAXPROCS(2)
	go leraning()
	fmt.Println("main")
}
