package main

import "fmt"

/*
	- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
func main() {

	n1 := struct{
		name string
		flavor string
		where []string
		vaiBemCom map[string]string

	}{
		name: "sorvete",
		flavor: "morango",
		where: []string{"Brasil", "EUA"},
		vaiBemCom: map[string]string{
			"Depois do almoço": "sobremesa",
			"a tarde": "com bolo",
		},
	}

	fmt.Println(n1)


}