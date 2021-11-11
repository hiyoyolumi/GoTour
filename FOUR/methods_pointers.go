package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//接收者
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//指针接收者
func (v *Vertex) Scale(f float64) { //不加指针的话，修改的是值的副本
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := &Vertex{3, 4}
	// v := Vertex{3, 4}
	v.Scale(10)
	//(&v).Scale(10)
	fmt.Println(v.Abs())
}
