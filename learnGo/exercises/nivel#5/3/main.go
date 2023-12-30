package main

import "fmt"

/*
- Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
*/

type Veiculo struct {
	portas int
	cor string
}
type Caminhonete struct {
	Veiculo
	traçãoNasQuatro bool

}
type Sedan struct {
	Veiculo
	modeloLuxo bool
}
func main() {

	c1 := Caminhonete{
		Veiculo: Veiculo{
			portas: 4,
			cor: "Prata",
		},
		traçãoNasQuatro: true,
	}
	fmt.Println(c1)

	c2 := Sedan{
		Veiculo: Veiculo{
			portas: 2,
			cor: "branco",
		},
		modeloLuxo: false,
	}
	fmt.Println(c2) 

	fmt.Println("Este carro possui tração nas 4 rodas ? ",c1.traçãoNasQuatro)
	fmt.Println("Este carro é um modelo de luxo ? ",c2.modeloLuxo)
}