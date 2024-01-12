package main

/*
	- Alem da goroutine principal, crie duas outras goroutines.
	- Cada goroutine adicional devem fazer um print separado.
	- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	novasgoroutines(100)
	wg.Wait()
}

func novasgoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine número:", i+1)
			wg.Done()
		}(x)
	}
}
