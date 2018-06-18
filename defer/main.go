package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println()

	answer()

	fmt.Println()
	fmt.Println()

	file("test")
	file("bad")
}

func answer() {
	fmt.Print("4")
	fmt.Print("2")
}

func open(name string) (*os.File, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func file(name string) {
	f, err := open(name)
	if err != nil {
		fmt.Printf("file '%s' can't be opened: %v\n", name, err)
		return
	}
	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	fmt.Println(string(b))
}
