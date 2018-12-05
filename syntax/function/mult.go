package main

import (
	"fmt"
)

func mult(a, b string) (string, string) {
	return a, b
}

func main() {
	x, y := mult("Hello", "World")
	fmt.Println(x, y)
}
