package main

import "fmt"

func retornaFuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}

func main() {

	x := retornaFuncao()
	y := x(3)

	fmt.Println(y)
	fmt.Println(retornaFuncao()(6))
}