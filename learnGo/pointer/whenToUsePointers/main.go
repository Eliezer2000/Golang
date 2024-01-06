package main

import "fmt"

func main() {

	x := 0
	estabeleceValor(x)
	estabelecePonteiro(&x)

	fmt.Println(x)
}

func estabeleceValor(x int) {
	x++
	fmt.Println("Na função: ", x)
}

func estabelecePonteiro(x *int){
	*x++
	fmt.Println("")
}