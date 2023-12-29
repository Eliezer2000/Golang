package main

import "fmt"

/*
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map
inteiro utilizando range.
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
	fmt.Println("\n=============================================================\n")

	delete(ck, "ronaldo_cristiano")

	fmt.Println("\n=============================================================\n")
	for k, v := range ck{
		fmt.Println(k)
		for i, h := range v{
			fmt.Println("\t", i, " - ", h)
		}
	}
}