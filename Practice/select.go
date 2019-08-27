// package main

// import "fmt"

// func fib(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("close")
// 			return
// 		default:
// 			fmt.Println("im in default")
// 		}
// 	}
// }
// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func() { //go must be function call
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}() // it is a function call
// 	fib(c, quit)
// }

package main

import "time"
import "fmt"

func main() {
	//For our example we’ll select across two channels.

	c1 := make(chan string)
	quit := make(chan string)
	//Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines.

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		quit <- "hello"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			return
			// case msg2 := <-c2:
			//     fmt.Println("received", msg2)
		case msg2 := <-quit:
			fmt.Println("close", msg2)
			//return
			//default:
			//	fmt.Println("im in default")
		}
	}
}
