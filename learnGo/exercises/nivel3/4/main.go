package main

import "fmt"

/*
	- Crie um loop utilizando a sintaxe: for {}
	- Utilize-o para demonstrar os anos desde que você nasceu.
*/

func main() {

	nasc := 2000
	anoAtual := 2024

	for  {
		if nasc > anoAtual {
			break
		}
		fmt.Println(nasc)
		nasc++
	}
}