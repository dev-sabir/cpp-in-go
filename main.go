package main

import (
	"ffi/swig/adder"
	"fmt"
)

func main() {
	a := adder.NewAdder()
	defer adder.DeleteAdder(a)

	a.Add(5)
  a.Add(20)
	fmt.Println("sum: ", a.Total())
}
