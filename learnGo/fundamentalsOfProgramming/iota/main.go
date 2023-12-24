package main

import "fmt"

/*
	- Numa declaração de constantes, o identificador iota representa números sequenciais.
- Na prática.
    - iota, iota + 1, a = iota b c, reinicia em cada const, _
*/

const (
	a = iota
	_ = iota
	x = iota
	z = iota
)

func main() {
	fmt.Println(a, x, z)
}
