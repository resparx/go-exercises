package main

import (
	"fmt"
	"learnpackage/add"
)

func main() {
	var a int = 1
	var b = 2
	c := "detox"
	fmt.Println("variables are", a, b, c)
	addd := add.Add(2, 3)

	fmt.Println("add func", addd)
}
