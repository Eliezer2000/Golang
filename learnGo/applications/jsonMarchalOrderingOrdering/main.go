package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	ContaBancaria float64
}

func main() {

	jamesBond := Pessoa{
		Nome:          "Joel",
		Sobrenome:     "silva",
		Idade:         55,
		Profissao:     "mecanico",
		ContaBancaria: 10.000,
	}

	blade := Pessoa{
		Nome:          "blade",
		Sobrenome:     "alves",
		Idade:         44,
		Profissao:     "caçador",
		ContaBancaria: 2.70,
	}
	j , err := json.Marshal(jamesBond)
	if err != nil {
		fmt.Println(err)
	}
	d, err := json.Marshal(blade)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(string(d))
}