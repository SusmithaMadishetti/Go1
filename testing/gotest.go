package gotest

import "errors"

//Div is a func
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisor can not be 0")
	}
	return a / b, nil
}
