package main

import "fmt"

/*
	- Crie e utilize uma função anônima.
*/

func main() {

	x := 21
	func(X int){
		fmt.Println(x, "vezes 6 : ")
		fmt.Println(x * 6)
	}(x)

	func() {
		fmt.Println("Hello world")
	}()
}