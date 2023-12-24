package main

import "fmt"

/*
	- Crie uma variável de tipo string utilizando uma raw string literal.
	- Demonstre-a.
*/

func main() {
	x := `Hello
	
	world 
	!`

	fmt.Println(x)
}