package main

import "fmt"

/*
	- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	- Demonstre estes valores.
*/

const (
	_ = 2000 + iota
	b
	c
	d
	e
)
func main() {
	fmt.Println(b, c, d, e)
}