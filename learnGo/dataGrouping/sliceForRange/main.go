package main

import "fmt"

/*
- Slices:
    - Tamanho: len(x)
    - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}
*/
 

func main() {
	slice := []int{20, 21, 22, 23, 1, 17}

	total := 0

	for _, valor := range slice {

		total += valor

	}
	fmt.Println("O valor total é: ", total)
}