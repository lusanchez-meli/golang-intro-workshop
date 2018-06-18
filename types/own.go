// :bn for next file
package main

import "fmt"

type ID int

func (i ID) Valid() bool {
	return i > 0
}

func main() {
	fmt.Println()

	ids := []ID{1, -3, 0}

	for _, id := range ids {
		fmt.Printf("%v is valid? %t\n", id, id.Valid())
	}

	var id ID = 5
	fmt.Printf("%v is valid? %t\n", id, id.Valid())
}
