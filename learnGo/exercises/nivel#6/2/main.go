package main

import "fmt"

/*
	- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de
	todos os ints recebidos;
	- Passe um valor do tipo slice of int como argumento para a função;
	- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de
	todos os elementos da slice;
	- Passe um valor do tipo slice of int como argumento para a função.
*/
func main() {

	x := []int{2, 7, 9}
	fmt.Println(soma(x...))

	y := []int{2, 7, 9}
	fmt.Println(somaS(y))
}

func soma(x ...int) int {
	tot := 0
	for _, v := range x{
		tot += v
	}
	return tot
}

func somaS(x []int) int {
	tot := 0
	for _, v := range x{
		tot += v
	}
	return tot
}