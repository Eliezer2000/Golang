package main

import "fmt"

/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja
uma variável do tipo string com identificador "esporteFavorito".
*/

func main() {
	
	esporteFavorito := "futebol"

	switch esporteFavorito{
	case "futebol":
		fmt.Println("Gol!")
	case "volei":
		fmt.Println("Saque ponto")
	default:
		fmt.Println("ops")
	}
}