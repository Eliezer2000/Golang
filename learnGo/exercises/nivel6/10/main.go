package main

import "fmt"

/*
	- Demonstre o funcionamento de um closure.
	- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de
	uma variável alem de seu scope.
*/

func main() {
	x := i()
	fmt.Println(x())
	fmt.Println(x())

	
	b := i()
	fmt.Println(b())
	fmt.Println(b())
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}