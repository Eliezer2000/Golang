package main

import "fmt"

func inverteSinal(numero int) int {
	return numero * -1
}

func inverteSinalPonteiros( numero *int) {
	*numero = *numero * - 1
}

func main() {
	numero := 20
	numeroInvertido := inverteSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novoNumero := 10
	fmt.Println(novoNumero)
	inverteSinalPonteiros(&novoNumero)
	fmt.Println(novoNumero)
}