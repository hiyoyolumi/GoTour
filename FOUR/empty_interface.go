package main

import "fmt"

type Put struct {
	key   int
	value int
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = Put{
		key:   1,
		value: 2,
	}
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
