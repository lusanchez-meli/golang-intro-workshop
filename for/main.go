package main

import "fmt"

func main() {

	l := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for i := 0; i < 10; i++ {
		show(i, l[i])
	}
	fmt.Println()

	i := 0
	for i < 10 {
		show(i, l[i])
		i++
	}
	fmt.Println()

	for i := range l {
		show(i, l[i])
	}
	fmt.Println()

	for i, v := range l {
		show(i, v)
	}
	fmt.Println()

	m := map[string]int{"x": 42, "y": 1337}

	for k, v := range m {
		show(k, v)
	}
	fmt.Println()

	fmt.Println("before forever")
label:
	for {
		break
		continue
		goto
	}
	fmt.Println("after forever")

}

func show(k, v interface{}) {
	fmt.Printf("%v:%v ", k, v)
}
