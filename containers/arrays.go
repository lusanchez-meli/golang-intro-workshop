// :bn for next file
package main

import "fmt"

func main() {
	fmt.Println()

	a := [5]string{"A", "B", "C", "D", "E"}
	fmt.Println("a:", a, len(a), cap(a))

	fmt.Println()

	b := []string{"A", "B"}
	fmt.Println("b:", b, len(b), cap(b))

	b = append(b, "D", "E")
	fmt.Println("b:", b, len(b), cap(b))

	fmt.Println()

	c := a[1:4]
	fmt.Println("c:", c, len(c), cap(c))

	c[2] = "XXX"

	fmt.Println("a:", a, len(a), cap(a))
	fmt.Println("c:", c, len(c), cap(c))

	fmt.Println()

	var d []int
	if d == nil {
		fmt.Println("d is nil")
	}
}
