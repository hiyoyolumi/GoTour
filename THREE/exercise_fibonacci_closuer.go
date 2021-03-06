package main

import (
	"fmt"
)

func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		temp := a
		a = b
		b = temp + b
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", f())
	}
	println()
}
