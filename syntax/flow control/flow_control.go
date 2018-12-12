package main

import (
	"fmt"
	"time"
)

func main() {
	/*1. loop*/
	x := 1
	for x <= 5 {
		fmt.Println(x)
		x++
	}

	//loop sequence is same as above
	for y := 1; y <= 5; y++ {
		if y%2 == 0 {
			fmt.Println(y, "Even Number")
		} else {
			fmt.Println(y, "Odd Number")
		}
	}

	/*2. switch*/

	//type switch
	z := 4
	switch z {
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}

	//expression switch
	length := 9
	switch {
	case length <= 7:
		fmt.Println("Short")
	case length <= 8:
		fmt.Println("Normal")
	case length > 8:
		fmt.Println("Long")
	}

	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("it's a weekday")
	default:
		fmt.Println("it's the weekend")
	}
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("afternoon")
	}

	/*3. defer
	A defer statement defers the execution of a function until the surrounding function returns.
	The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	*/
	hello()
	count()

	/* used in open & close file:
	f,err := os.Open("home/...")
	if err !=nil {
		return err
	}
	defer f.Close()
	*/

	panicTest()
}

func hello() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func count() {
	fmt.Println("counting")

	for i := 0; i < 5; i++ {
		//arguments get evaluated right away, but the call is deferred
		//defer call will be put on a stack
		defer fmt.Println(i)
	}
	fmt.Println("finished")
}

func panicTest() {
	//unnamed function with the call to recover
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
