package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Diff() float64 {
	return v.X - v.Y
}

func Plus(v Vertex) float64 {
	return v.X + v.Y
}

func (f MyFloat) UnsignedAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{134.1114, 24.411411}
	fmt.Println(v.Abs())
	fmt.Println(v.Diff())
	fmt.Println(Plus(v))
	f := MyFloat(-math.SqrtPi)
	fmt.Println(f.UnsignedAbs())
}
