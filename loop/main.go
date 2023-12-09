package main

import (
	"fmt"
)

func main() {
	// loopFor()
	loopRange()
	// loopWhile()
}

func loopFor() {
	a := 3
	for i := 0; i < a; i++ {
		fmt.Println(i, " ")
	}
}

func loopRange() {
	a := []int{0, 1, 1, 2, 3, 5, 8}
	for _, i := range a {
		fmt.Println(i, " ")
	}
}

func loopWhile() {
	a := 3
	for a > 0 {
		fmt.Println(a, " ")
		a--
	}
}
