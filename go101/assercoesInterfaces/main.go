package main

import "fmt"

type Document interface {
	Doc()string
}
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

func(pf PessoaFisica) Doc() string {
	return fmt.Sprintf("Meu Cpf é %s ", pf.cpf)
}
type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	Cnpj string
}

func(pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("Meu Cnpj é %s ", pj.Cnpj)
}
func Show(d Document) {
	switch d.(type) {
	case PessoaFisica:
		fmt.Println(d.(PessoaFisica).Sobrenome)
	case PessoaJuridica:
		fmt.Println(d.(PessoaJuridica).RazaoSocial)
	}
	fmt.Println(d)
	fmt.Println(d.Doc())
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
	Show(p)

	pj := PessoaJuridica {
		Pessoa: Pessoa{
			Nome: "Jamal",
			idade: 30,
			Status: true,
		},
		RazaoSocial: "zzz",
		Cnpj: "000000.00",
	}
	Show(pj)
}