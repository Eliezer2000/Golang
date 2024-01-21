package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
	- Utilize atomic para consertar a condição de corrida do exercício #3.
*/

var w sync.WaitGroup

const qtdeGoroutines = 5

var contador int32

func main() {
	criarGoRoutines(qtdeGoroutines)
	w.Wait()
	fmt.Println("Total de Goroutines: ", qtdeGoroutines, "\nTotal do contador: ", contador)

}

func criarGoRoutines(n int) {
	w.Add(n)
	for i := 0; i < n; i++{
		go func() {
			atomic.AddInt32(&contador, 1)
			runtime.Gosched()
			v := atomic.LoadInt32(&contador)
			fmt.Println(v)
			w.Done()
		}()
	}
}