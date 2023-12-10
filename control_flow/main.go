package main

import (
	"fmt"
)

func main() {
	defer ifElse()
	switchCase()
}

func ifElse() {
	a := 5
	if a > 3 {
		a = a - 3
	} else if a == 3 {
		a = 0
	} else {
		a = a + 3
	}
	fmt.Println(a)
}

func switchCase() {
	a := "NO"
	switch a {
	case "YES":
		a = "Y"
	case "NO":
		a = "N"
	default:
		a = "X"
	}
	fmt.Println(a)
}
