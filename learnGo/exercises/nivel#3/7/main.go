package main

import "fmt"

// - Utilizando a solução anterior, adicione as opções else if e else.

func main() {

	on := false
	pause := false

	if on {
		fmt.Println("Open")
	} else if pause {
		fmt.Println("coffe")
	}else {
		fmt.Println("Close")
	}

}