package main

import "fmt"

/*
	   - Crie um tipo para um struct chamado "pessoa"
    - Crie um método "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
    - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
    - Crie uma função "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
    - Demonstre no seu código:
        - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
        - Que você não pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/

type pessoa struct {
	nome string
}
func (p *pessoa) falar() {
	fmt.Println("Olá me chamo: ", p.nome)
}
type humanos interface {
	falar()
}
func dizerAlgumaCoisa(h humanos){
	h.falar()
}
func main() {
	p1 := pessoa {
		nome: "joel",
	}

	dizerAlgumaCoisa(&p1)
}