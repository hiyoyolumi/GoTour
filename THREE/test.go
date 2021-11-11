package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	c := strings.Fields(s)
	for _, v := range c {
		m[v] += 1
	}
	return m
}

func main() {
	// fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	c := WordCount("haha hehe bingbing bb ddyyqq dyq dyq")
	fmt.Println(c)
}
