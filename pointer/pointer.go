package main

import "fmt"

func num(yPtr *int) {
	*yPtr = 3
}

func main() {
	x := 7
	num(&x)
	fmt.Println(x)
	yPtr := new(int)
	num(yPtr)
	fmt.Println(*yPtr)
}
