package main

import (
	"fmt"
	"time"
)

func numeros(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel %d\n", i)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Fim da escrita!")
	close(n)
}


func main(){
	cn := make(chan int)
	go numeros(cn)
	for v := range cn {
		fmt.Printf("Lido do channel %d\n", v)
	}
	
	fmt.Println("Fim da execução")
}
