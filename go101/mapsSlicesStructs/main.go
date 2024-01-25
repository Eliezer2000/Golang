package main

import "fmt"

//STRUCT
type Pessoa struct {
	Nome string
	Sobrenome string
	idade uint8
	Status bool
	cpf string
}

func main() {
	nomes := []string{"Joel", "Joana", "Jamal"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Julia")
	fmt.Println(nomes, len(nomes), cap(nomes))


	//MAP
	idades := make(map[string]uint8)
	idades["Joel"]  = 30
	idades["Joana"] = 26
	idades["Jamal"] = 44
	fmt.Println(idades)

	val, ok := idades["Joao"]
	fmt.Println(val, ok)


	//STRUCT
	p := Pessoa {
		Nome: "Joel",
		Sobrenome: "Silva",
		idade: 25,
		Status: true,
		cpf: "123344",
	}
	fmt.Println(p)
}