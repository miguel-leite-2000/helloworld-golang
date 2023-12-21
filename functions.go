package main

import "fmt"

type Car struct {
  Name string
}

func(c Car) toWalk() {
  fmt.Println(c.Name, "Walked")
}
// main is the entry point for the application.
func main() {

  car := Car {
    Name: "BMW",
  }

  car.toWalk()

  sumAllNumbers := func (x ... int) func() int {
    res := 0
    for _, v := range x {
      res += v
    }
    return func() int {
      return res * res
    }
  }

	// result := sumAll(10, 20,10, 20,10, 20, 10, 20)

	fmt.Println(sumAllNumbers(10, 20,10, 20,10, 20, 10, 20)())
}

// Sum returns the sum of two integers.
func Sum(a int, b int) (result int) {
	result = a + b
	return
}

// sumAll returns the sum of all its arguments.
// It panics if there are any duplicates in the argument list.
func sumAll(x ...int) int {
  result := 0
  for _, v := range x {
    result += v
  }
  return result
}