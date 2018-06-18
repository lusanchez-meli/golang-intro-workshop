// :bn for next file
package main

import (
	"fmt"
)

func pow2(x int) int {
	return x * x
}

func main() {
	fmt.Println()

	n := pow2(3)
	fmt.Println(n)

	fmt.Println()

	l := []int{2, 4, 6, 8}

	fmt.Print("    for: ")
	for _, v := range l {
		fmt.Print(pow2(v), " ")
	}
	fmt.Println()

	fmt.Print(" mapper: ")
	fmt.Println(mapper(pow2, l))
}

func mapper(fn func(int) int, list []int) []int {
	r := make([]int, len(list), cap(list))
	for i, v := range list {
		r[i] = fn(v)
	}
	return r
}
