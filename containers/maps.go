package main

import "fmt"

type Point struct {
	X, Y float64
}

func main() {
	fmt.Println()

	m := map[string]int{"hello": 5, "world": 42}
	fmt.Println(m)

	ps := map[string]Point{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(ps)

	fmt.Println()

	v := m["hello"]
	fmt.Println(v)

	w, ok := m["world"]
	fmt.Println(w, ok)

	bad, ok := m["bad"]
	fmt.Println(bad, ok)

	fmt.Println()

	delete(m, "hello")
	fmt.Println(m)
}
