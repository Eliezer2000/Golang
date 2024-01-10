package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	var contador int64
	contador = 0
	totalgoroutines := 10

	var wg sync.WaitGroup
	wg.Add(totalgoroutines)


	for i := 0; i < totalgoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador: \t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Valor final: ", contador)
}