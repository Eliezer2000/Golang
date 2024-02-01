package main

import "fmt"

var n int

func main() {
	fmt.Println("Função Main sendo executada")
	n = 10
}

func init() {
	fmt.Println("Executando a função init")
	fmt.Println(n)
}