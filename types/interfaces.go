package main

import (
	"fmt"
	"time"
)

type Stringer interface {
	String() string
}

type ID time.Time

func (i ID) String() string {
	return fmt.Sprintf("%d", i)
}

func (i ID) Valid() bool {
	return i > 0
}

type Name string

func (n Name) String() string {
	return string(n)
}

func main() {
	fmt.Println()

	ss := []Stringer{
		ID(1),
		ID(-3),
		Name("Juan"),
		ID(0),
		Name("Carlos"),
	}

	// for _, s := range ss {
	// 	fmt.Println(s.String())
	// }

	// fmt.Println()

	// id, ok := ss[2].(ID)
	// fmt.Println("id:", id)

	for _, s := range ss {
		id, ok := s.(ID)
		if !ok {
			fmt.Printf("'%v' is not an ID\n", s)
		} else {
			fmt.Println("ID:", id)
		}
	}

}
