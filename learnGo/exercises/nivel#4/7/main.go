package main

import "fmt"

/*
	- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
	- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/

func main() {

	ss := [][]string {
		[]string {
			"joel",
			"silva",
			"filmes e series",
		},
		[]string {
			"Maria", 
			"gomes",
			"cozinhar",
		},
		[]string {
			"Jose",
			"rocha",
			"futebol",
		},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	fmt.Println("\n=============================================================\n")

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v{
			fmt.Println("\t", item)
		}
	}

}