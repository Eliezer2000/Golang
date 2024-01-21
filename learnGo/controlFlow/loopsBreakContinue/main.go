package main

import "fmt"

/*
	- Operação módulo: %
	- For: break
	- For: continue
*/

func main() {

	for i := 0; i < 20; i++ {
		if i % 2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}