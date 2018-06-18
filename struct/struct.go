// :bn for next file
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Distance() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) String() string {
	return fmt.Sprintf("(%v, %v) %v", p.X, p.Y, p.Distance())
}

func main() {
	fmt.Println()

	v := Point{X: 3, Y: 4} // type Point
	fmt.Println("v:", v)

	z := Point{} // type Point
	fmt.Println("z:", z)

	p := &v // type *Point
	fmt.Println("p:", p)

	v.X = 0
	fmt.Println("v:", v)

	fmt.Println("p:", p)
}
