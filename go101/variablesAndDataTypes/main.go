package main

import "fmt"

var (
	name string
	n1   int
	n2   int
)

func main() {
	//declara e inicializa
	mensagem := "Hello world"
	fmt.Println(mensagem)

	var a, b, c = true, 2.3, "hello"
	fmt.Println(a, b, c)

	var total int
	total++
	fmt.Println("total: ", total)

	name = "joel"
	fmt.Println(name)

	var x, y = 10, 20
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y)
}