package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scala(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScalaFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scala(3)
	ScalaFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scala(3)
	ScalaFunc(p, 8)

	fmt.Println(v, p)
}
