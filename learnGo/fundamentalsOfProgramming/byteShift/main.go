package main

import "fmt"

//- Deslocamento de bits é quando deslocamos digitos binários para a esquerda ou direita.

func main() {
	x := 1
	y := x >> 10
	fmt.Printf("%d\n", x)
	fmt.Printf("%d\n", y)
}
