package main

import "fmt"

type usuario struct {
	nome  string
	idade int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua A", 0}
	
	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Davi"}
	fmt.Println(usuario3)
}