package main

import "fmt"

var (
	f1, n, f2, temp int
)

func fb(n int) {
	f1 := 0
	f2 := 1
	fmt.Print(f1)
	fmt.Print(f2)
	for i := 0; i < n-2; i++ {

		temp = f1 + f2
		f1 = f2
		f2 = temp
		fmt.Print(temp)

	}

}
func fi(i int) int {
	if i <= 1 {
		return i
	}
	return fi(i-1) + fi(i-2)

}
func main() {
	fmt.Println("enter number")
	fmt.Scanf("%d", &n)
	fb(n)
	for i := 0; i < n; i++ {
		fmt.Print(fi(i))

	}

}
