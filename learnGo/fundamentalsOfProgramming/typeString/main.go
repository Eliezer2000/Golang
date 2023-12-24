package main

import "fmt"

/*
- Strings são sequencias de bytes.
- Imutáveis.
- Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
- Na prática:
    - %v %T
    - Raw string literals
    - Conversão para slice of bytes: []byte(x)
    - %#U, %#x
*/
func main() {
	s := "Hello world"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T %#U - %#X\n", v, v, v, v)
	}

	fmt.Printf("%v, %T", sb, sb)
}
