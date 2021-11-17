package main

import "fmt"

type Vertex struct {
	a, b int
}

func (m *Vertex) Request() *Vertex {
	if m != nil {
		m.a = 10
		m.b = 20
		return nil
		// return &Vertex{
		// 	a: 10,
		// 	b: 20,
		// }
	}
	return nil
}

func main() {
	res := new(Vertex)
	res.a = 1
	res.b = 2
	fmt.Println(res)
	// res = res.Request()
	res.Request()
	fmt.Println(res)
}
