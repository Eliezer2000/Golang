package main

import "fmt"

func main() {
	nomes := []string{"Joel", "Joana", "Jamal"}
	for k, nome := range nomes {
		fmt.Println(k, nome)
	}

	var i int
	for i < len(nomes) {
		fmt.Println(nomes[i])
		i++
	}
}