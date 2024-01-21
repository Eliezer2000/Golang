package main

import "fmt"

/*
	- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
	- Demonstre todos esses valores e seus índices.
*/

func main() {

	ck := map[string][]string{
		"Wayne_Bruce": []string{
			"sair a noite", "carros luxuosos",
		},
		"senna_ayrton": []string{
			"correr", "praicar esportes",
		},
	}
	fmt.Println(ck)

	fmt.Println("\n=============================================================\n")

	for k, v := range ck{
		fmt.Println(k)
		for i, h := range v{
			fmt.Println("\t", i, " - ", h)
		}
	}
}