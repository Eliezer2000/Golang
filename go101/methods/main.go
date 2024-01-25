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

func(p Pessoa) String() string {
	return fmt.Sprintf("Olá meu nome é %s e eu tenho %d anos", p.Nome, p.idade)
}
type Categoria struct {
	Nome string
	Pai *Categoria
}

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

func (c *Categoria) SetPai(pai *Categoria) {
	c.Pai = pai
}
func main() {
	//STRUCT
	p := Pessoa {
		Nome: "Joel",
		Sobrenome: "Silva",
		idade: 25,
		Status: true,
		cpf: "123344",
	}
	fmt.Println(p)

	cat := Categoria {
		Nome: "categoria 1",
	}
	cat.SetPai(&Categoria{Nome: "Pai"})
	if !cat.HasParent() {
		fmt.Println("Foi comprar cerveja")
	}
	fmt.Println(cat)
}