package main

import (
	"fmt"
)

type Hero struct {
	name string
	role string
}

/*
Method with value receiver
*/
func (e Hero) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Hero) changeRole(role string) {
	e.role = "Feeder"
}

func main() {
	e := Hero{
		name: "Linglung",
		role: "Assasin",
	}
	fmt.Printf("Hero name before change: %s\n", e.name)
	e.changeName("Hayabuset")
	fmt.Printf("Hero name after change: %s\n\n", e.name)
	
	fmt.Printf("Hero role before change: %s\n", e.role)
	(&e).changeRole("Feeder")
	fmt.Printf("Hero role after change: %s\n", e.role)
}
