package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	- Utilize mutex para consertar a condição de corrida do exercício anterior.
*/

var w sync.WaitGroup

const qtdeGoroutines = 5

var contador int

var mu sync.Mutex

func main() {
	criarGoRoutines(qtdeGoroutines)
	w.Wait()
	fmt.Println("Total de Goroutines: ", qtdeGoroutines, "\nTotal do contador: ", contador)

}

func criarGoRoutines(n int) {
	w.Add(n)
	for i := 0; i < n; i++{
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			w.Done()
		}()
	}
}