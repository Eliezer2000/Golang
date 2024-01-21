package main

import "fmt"

//- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

func main() {
	x := 100
	fmt.Printf("%d , %#x , %b", x, x, x)
}