package main

import (
	"fmt"
	"math"
)

/*
	- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"

*/
type quadrado struct{
	lado1 float64
}
type circulo struct{
	raio float64
}

func(q quadrado) area(){
	 res := q.lado1 * q.lado1
	 fmt.Println("Area do quadrado: ",res)
}

func(c circulo) area()  {
	res := math.Pi * 2 * c.raio
	fmt.Println("Area do circulo ", res)
}

type info interface{
	area()
}

func medida(i info) {
	i.area()
}

func main(){

	quad := quadrado{
		lado1: 15.8,
	}
	quad.area()

	circ := circulo{
		raio: 22.5,
	}
	circ.area()
	medida(quad)
	medida(circ)

}