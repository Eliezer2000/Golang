package main

import "fmt"

/*
	- For: inicialização, condição, pós
	- For: condição ("while")
	- For: ...ever? (http servers)
	- For: break
*/

func main() {
	x := 0

	for {
		if x < 10 {
			x++
			fmt.Println("x é menor que 10")
		}else{
			fmt.Println("x é maior que 10, STOP!!!")
			break
		}
	}
}