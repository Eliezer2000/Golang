package main

import "fmt"

/*
	- Structs dentro de structs dentro de structs.
	- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) e
	tambem um competidor (nome, equipe, pontos).
*/
type Person struct{
	name string
	age int
}

type Professional struct {
	Person
	title string
	salario int
}

func main() {
	p1 := Person{
		"Joel", 
		30,
	}

	p2 := Professional{
		Person: Person{
			name:"Ana", 
			age: 25,
		},
		title: "teacher",
		salario: 5679,
	}

	p3 := Person{"Marcel", 44}
	P4 := Professional{Person{"JOhn", 80}, "Jogador", 50}
	
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.age)
	fmt.Println(p1.age)
	fmt.Println(p3)
	fmt.Println(P4)
	
}