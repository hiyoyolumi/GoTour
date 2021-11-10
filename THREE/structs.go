package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var p Vertex = Vertex{3, 4}
	// p = Vertex{
	// 	1, 2,
	// }
	fmt.Println(p)
	// fmt.Println(Vertex{1, 2})
}
