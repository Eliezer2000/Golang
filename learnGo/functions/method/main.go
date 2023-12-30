package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) bomdia() {
	fmt.Println(p.nome, " Bom dia!")
}
func main() {

	joel := Pessoa{"Joel", 55}
	joel.bomdia()
}