package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot sqrt negative number : %g", e)
}

func Sqrt(x float64) (float64, error) {
	res := 1.0
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		for n := 0; n < 10; n++ {
			res = res - ((res*res - x) / (2 * res))
		}
		return res, nil
	}
}
func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
}
