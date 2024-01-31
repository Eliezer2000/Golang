package main

import "fmt"

func main() {
	nomes := [3]string{"João", "Davi", "Joel"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"Nome": "Joel",
		"Sobrenome": "Silva",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}