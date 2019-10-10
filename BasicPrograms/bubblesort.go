package main

import "fmt"

func main() {
	num := []int{34, 2, 5, 9, 34, 45, 45, 3, 4}
	fmt.Println(bubblesort(num))
}
func bubblesort(num []int) []int {
	for i := len(num); i > 0; i-- {
		for j := 1; j < i; j++ {
			if num[j-1] > num[j] {
				temp := num[j]
				num[j] = num[j-1]
				num[j-1] = temp
			}
		}

	}
	return num
}
