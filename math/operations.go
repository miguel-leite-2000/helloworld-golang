package math

import "errors"

var A string = "SHOWWW"

// Sum is a function that adds the values ​​entered in the function parameters and returns int
func Sum(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("The result is greater than 10")
	}

  return res, nil;
}