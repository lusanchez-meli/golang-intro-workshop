package main

import "fmt"

func main() {

	fmt.Println()

	if true {
		fmt.Println("1")
	}

	if false {
		fmt.Println("no")
	} else {
		fmt.Println("2")
	}

	if false {
		fmt.Println("no")
	} else if false {
		fmt.Println("neither")
	} else {
		fmt.Println("3")
	}

	if v := fn(); v > 0 {
		fmt.Println("4")
	}

}

func fn() int {
	return 42
}
