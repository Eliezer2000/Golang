package main

import "fmt"

/*
	- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro
	utilizando range.
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

	ck["ronaldo_cristiano"] = []string{"treinar", "ficar com a familia"}

	for k, v := range ck{
		fmt.Println(k)
		for i, h := range v{
			fmt.Println("\t", i, " - ", h)
		}
	}
}