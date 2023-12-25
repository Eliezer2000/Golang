package main

import "fmt"

/*
	- Switch:
    - pode avaliar uma expressão
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos
*/

func main() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")
	}

	y := 9
	switch true {
	case (y == 4), (y == 8):
		fmt.Println("4 ou 8")
	case(y < 10):
		fmt.Println("3 ou 4")
	default:
		// Padrão casa nao houver outra opção
	}
}