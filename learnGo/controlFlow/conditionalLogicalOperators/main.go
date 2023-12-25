package main

import "fmt"

/*
	- &&
	- ||
	- !

	- true && true
    - true && false
    - true || true
    - true || false
    - !true
*/

func main() {
	x := 9

	if !(x % 2 == 0) && x % 3 == 0 {
		fmt.Println("É multiplo de 2 e tambem de 3")
	}
	if x % 2 == 0 || x % 3 == 0 {
		fmt.Println("É multiplo de 2 ou 3")
	}
}