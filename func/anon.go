package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println()

	n := add(1, 2) // type: func(int, int) int
	fmt.Println(n)

	add5 := curry(add, 5) // type: func(int) int
	m := add5(2)
	fmt.Println(m)

	add2 := curry2(add, 2) // type: main.IntToInt
	o := add2(40)
	fmt.Println(o)

	fmt.Println()

	fmt.Printf("  add: %T\n", add)
	fmt.Printf(" add5: %T\n", add5)
	fmt.Printf(" add2: %T\n", add2)

}

func curry(fn func(int, int) int, v int) func(int) int {
	return func(y int) int {
		return fn(v, y)
	}
}

type IntIntToInt func(int, int) int
type IntToInt func(int) int

func curry2(fn IntIntToInt, v int) IntToInt {
	return func(y int) int {
		return fn(v, y)
	}
}
