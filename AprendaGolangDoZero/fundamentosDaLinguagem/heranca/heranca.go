package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"Joao", "Silva", 26, 188}
	fmt.Println(p1)
	e1 := estudante{p1, "Programação", "USP"}
	fmt.Println(e1)
}