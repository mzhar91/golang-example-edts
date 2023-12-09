package main

import "fmt"

func main() {
	// variable()
	// types()
	ml()
}

func variable() {
	// manually specifying type to the variable
	var first string
	first = "This is first string"
	
	// define type to the variable
	var second = "This is second string"
	
	// define type to the variable
	third := "This is third string"
	
	fmt.Println(first, second, third)
}

func types() {
	a := 1    // var a int
	b := 3.14 // var b float
	c := "hi" // var c string
	d := true // var d bool
	fmt.Println(a, b, c, d)
	
	e := []int{1, 2, 3} // slice
	e = append(e, 4)
	fmt.Println(e, len(e), e[0], e[1:3], e[2:], e[:2])
	
	f := make(map[string]int) // map
	f["one"] = 1
	f["two"] = 2
	fmt.Println(f, len(f), f["one"], f["three"])
}

type Hero struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	isMeta bool
}

func ml() {
	p := Hero{
		Name:   "Linglung",
		Role:   "Assasin",
		isMeta: true,
	}
	fmt.Println(p, p.Name, p.Role, p.isMeta)
}

