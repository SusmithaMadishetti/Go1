package main

import "fmt"

func main() {
	a := []int{6, 3, 8, 6, 8, 2, 1}
	fmt.Println(bubble(a))
}
func bubble(num []int) []int {
	for i := len(num); i > 0; i-- {
		for j := 1; j < len(num); j++ {
			if num[j-1] > num[j] {
				//temp := num[j-1]
				// num[j-1] = num[j]
				// num[j] = temp
				num[j-1], num[j] = num[j], num[j-1]
			}
		}
	}
	return num
}
