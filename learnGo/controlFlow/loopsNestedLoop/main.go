package main

import "fmt"

/*
	- For
    - Repetição hierárquica
    - Exemplos: relógio, calendário
*/
func main(){

	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Horas: ", horas)
		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}
		println("\n")
	}

}