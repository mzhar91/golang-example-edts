package main

import (
	"fmt"
)

func main() {
	// exampleMain()
	// getAddress()
	getValue()
}

func exampleMain() {
	x := 0
	y := 0
	fmt.Printf("before:\ta=%d\tb=%d\n", x, y)
	
	// call exampleUpdate
	exampleUpdate(x, &y)
	fmt.Printf("after:\ta=%d\tb=%d\n", x, y)
}

func exampleUpdate(a int, b *int) {
	a = 10
	*b = 10
	fmt.Printf("inside:\ta=%d\tb=%d\n", a, *b)
}

func getAddress() {
	a := 10
	b := &a
	fmt.Println(b)
}

func getValue() {
	a := 10
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
}
