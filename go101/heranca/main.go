package main

import "fmt"

type Pessoa struct {
	Nome      string
	idade     uint8
	Status    bool
}

func(p Pessoa) String() string {
	return fmt.Sprintf("Olá meu nome é %s e eu tenho %d anos", p.Nome, p.idade)
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

type PessoaJuridica struct {
	RazaoSocial string
	Cnpj string
}


func main() {
	p := PessoaFisica {
		Pessoa: Pessoa{
			Nome: "Joel",
			idade: 30,
			Status: true,
		},
		Sobrenome: "Silva",
		cpf: "000.000.000.00",
	}
	fmt.Println(p)

}