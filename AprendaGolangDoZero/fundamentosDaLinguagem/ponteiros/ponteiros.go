package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	variavel2++
	fmt.Println(variavel1,variavel2)

	// Ponteiro é uma referencia de memoria
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}