package main

import "fmt"

func main() {

	amigos := map[string]int{
		"Alfred": 123456,
		"Bruce":  654321,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["Bruce"])

	amigos["Pinguim"] = 44444

	fmt.Println(amigos)
	fmt.Println(amigos["Pinguim"])

	if sera, ok := amigos["Pinguim"]; !ok{
		fmt.Println("Não tem")
	} else {
		fmt.Println(sera)
	}

}