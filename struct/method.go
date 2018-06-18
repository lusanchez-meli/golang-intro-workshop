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

	fmt.Println(v.String())
	fmt.Println(Point.String(v))

	fmt.Println()

	fmt.Println(v.Distance())
	fmt.Println((*Point).Distance(&v))
}
