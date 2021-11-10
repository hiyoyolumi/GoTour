package main

import "strings"

func main() {
	a := "hahaha"
	b := "hehehe"
	c := strings.Join([]string{a, b}, ",")
	println(c)
}
