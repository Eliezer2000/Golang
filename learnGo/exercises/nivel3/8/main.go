package main

import "fmt"

/*
	- Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/

func main() {

	on := true
	pause := false

	switch{
	case on:
		fmt.Println("Open")
	case pause:
		fmt.Println("coffe")
	default:
		fmt.Println("Close")
	}


}