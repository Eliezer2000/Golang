package main

import (
	"fmt"
	
)

func main() {

	data := map[int]string{
		123: "nome",
		321: "idade",
		765: "pais", 
		987: "regiao",
	}
	fmt.Println(data)

	total := 0

	for key, _ := range data{
		total += key
	}
	fmt.Println(total)

	delete(data, 123)
	fmt.Println(data)
}