package main

import "fmt"

/*
	- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map,
	utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range
anterior.
*/

type Pessoa struct{
	nome string
	sobrenome string
	sabores []string
}

func main() {

	mapa := map[string]Pessoa{
		"Tijolo":Pessoa{
			nome: "joel",
			sobrenome: "silva",
			sabores: []string{"napoleao", "morango", "baunilha"},
		},

	}

	for _, v := range mapa {
		fmt.Println("Meu nome é :", v.nome, " e meus sorvetes favoritos são: ")
		for _, v := range v.sabores{
			fmt.Println(v)
		}
	}


}