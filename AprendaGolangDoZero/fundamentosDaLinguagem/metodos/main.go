package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados d usuario %s no database", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 20
}

func (u *usuario) fazerAniversario(){
	u.idade++
}
func main() {
	usuario1 := usuario{"Usuario", 21}
	fmt.Println(usuario1)
	usuario1.salvar()

	maior := usuario1.maiorIdade()
	fmt.Println(maior)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}