package main

import "fmt"

/*
	- São structs sem identificadores.
	- x := struct { name type }{ name: value }
*/

func main() {
	x := struct {
		name string
		age int
	}{
		name: "joel",
		age: 44,
	}

	fmt.Println(x)
}