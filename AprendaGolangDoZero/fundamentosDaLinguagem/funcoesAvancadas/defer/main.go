package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2(){
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada ")
	fmt.Println("Entrando na função para verificar se o aluno esta aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {

	fmt.Println(alunoAprovado(5, 5))
}