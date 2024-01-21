package main

import "fmt"

/*
	- Crie uma função que retorna uma função.
	- Atribua a função retornada a uma variável.
	- Chame a função retornada.
*/

func main(){
	x := RetornaFunc()
	x()
}

func RetornaFunc() func() {
	return func(){
		fmt.Println("Retornando função")
	}
	
}