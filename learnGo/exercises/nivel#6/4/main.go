package main

import "fmt"

/*
	- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.

*/

type Pessoa struct{
	nome string
	sobrenome string
	idade int
}
func(p Pessoa)dadosPessoal (){
	fmt.Println("Primeiro nome: ", p.nome, "\nSeu sobrenome:  ", p.sobrenome, "\nSua idade: ", p.idade)
}
func main() {

	x1 := Pessoa{
		nome: "Joel",
		sobrenome: "Silva",
		idade: 44,
	}
	Pessoa.dadosPessoal(x1)
}