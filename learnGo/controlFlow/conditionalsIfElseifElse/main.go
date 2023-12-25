package main

import "fmt"

/*
	- If, else.
	- If, else if, else.
	- If, else if, else if, ..., else.
*/

func main() {
	if x := 10; x > 100 {
		fmt.Println("x é maior que 100")
	} else if x < 100 {
		fmt.Println("x é menor que 100")
	}else {
		fmt.Println("x não é menor que 100 e nem maior que 100")
	}
}