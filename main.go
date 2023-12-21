package main

import (
	"fmt"
	"helloworld/math"
	"log"
)

func main() {

	res, err := math.Sum(10, 10)

  if err!= nil {
    log.Fatal(err.Error())
  }

  fmt.Println(res)
}
