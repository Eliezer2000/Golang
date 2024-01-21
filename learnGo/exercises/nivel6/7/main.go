package main

import "fmt"

/*
	- Atribua uma função a uma variável.
	- Chame essa função.
*/

func main() {
	x := func ()  {
		fmt.Println("Hello")
	}
	x()
}