package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	- Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
*/
var w sync.WaitGroup

const qtdeGoroutines = 5

var contador int

func main() {
	criarGoRoutines(qtdeGoroutines)
	w.Wait()
	fmt.Println("Total de Goroutines: ", qtdeGoroutines, "\nTotal do contador: ", contador)

}

func criarGoRoutines(n int) {
	w.Add(n)
	for i := 0; i < n; i++{
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			w.Done()
		}()
	}
}