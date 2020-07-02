package main

import (
	"fmt"
	"learnpackage/add"
)

func main() {
	var a int = 1
	var b = 2
	c := "detox"

	arr := [...]int{1, 2, 3, 4, 5, 6, 6, 7, 8}
	slArr := arr[3:5]
	fmt.Println("arr", arr)
	fmt.Println("slArr", slArr)
	fmt.Println("variables are", a, b, c)
	addd := add.Add(2, 3)

	fmt.Println("add func", addd)
}
