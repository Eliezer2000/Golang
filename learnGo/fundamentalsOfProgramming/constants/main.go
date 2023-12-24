package main

import "fmt"

/*
	- São valores imutáveis.
	- Podem ser tipadas ou não:
    - const oi = "Bom dia"
    - const oi string = "Bom dia"
	- As não tipadas só terão um tipo atribuido a elas quando forem usadas.
    - Ex. qual o tipo de 42? int? uint? float64?
    - Ou seja, é uma flexibilidade conveniente.
	- Na prática: int, float, string.
    - const x = y
    - const ( x = y )
*/

const (
	x = 10
	y = 20
	z = 30
)

var d float64

func main() {
	d = x
	fmt.Println(d)
	fmt.Println(x, y, z)
}
