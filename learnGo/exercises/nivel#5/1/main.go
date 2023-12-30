package main

import "fmt"

/*
- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice
	que contem os sabores de sorvete.
*/

type Pessoa struct{
	nome string
	sobrenome string
	sabores []string
}

func main() {

	p1 := Pessoa{
		nome: "joel",
		sobrenome: "silva",
		sabores: []string{"napoleao", "morango", "baunilha"},
	}

	p2 := Pessoa{nome:"maria", sobrenome: "rosa", sabores: []string{"chocolate", "cacau"}}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.nome, p1.sobrenome)
	for _, v := range p1.sabores {
		fmt.Println("\t ", v)
	}
	fmt.Println(p2.nome, p2.sobrenome)
	for _, v := range p2.sabores {
		fmt.Println("\t ", v)
	}
}