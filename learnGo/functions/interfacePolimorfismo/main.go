package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type Dentista struct {
	Pessoa
	dentesArrancados int
	salario          float64
}

type Arquiteto struct {
	Pessoa
	tipoConstrucao string
	tamanhoLocura  string
}

func (x Dentista) Bomdia() {
	fmt.Println("Olá meu nome é ", x.nome, " nessa semana arranquei ", x.dentesArrancados, " dentes.")
}

func (x Arquiteto) Bomdia(){
	fmt.Println("Meu nome é ", x.nome, " bora fazer um projeto ? ")
}

type gente interface {
	Bomdia()
}

func serhumano(g gente) {
	g.Bomdia()
}
func main() {
	DrArnaldo := Dentista{
		Pessoa: Pessoa{
			nome: "Arnaldo",
			sobrenome: "Foster",
			idade: 48,
		},
		dentesArrancados: 9,
		salario: 56000,
	}

	arq1 := Arquiteto{
		Pessoa: Pessoa{
			nome: "Ana",
			sobrenome: "Castro",
			idade: 34,
		},
		tipoConstrucao: "imperial",
		tamanhoLocura: "medio",
	}

	fmt.Println(DrArnaldo)
	fmt.Println(arq1)

	DrArnaldo.Bomdia()
	arq1.Bomdia()
}