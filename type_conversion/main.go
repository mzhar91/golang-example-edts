package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	conversion()
}

func conversion() {
	a := 3.14
	b := int(a)
	fmt.Println(b, reflect.TypeOf(b))
	
	c := "12.34"
	d, _ := strconv.ParseFloat(c, 64)
	fmt.Println(d, reflect.TypeOf(d))
	
	e := false
	f := fmt.Sprint(e)
	fmt.Println(f, reflect.TypeOf(f))
}
