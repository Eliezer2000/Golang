package main

import "fmt"

/*
- ...use um for range loop para demonstrar os valores do canal.
*/

func main(){
	c := gen()
	go receive(c)

	fmt.Println("About to exit")
}

func gen() <-chan int{
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++{
			c <- i
		}
		close(c)		
	}()
	
	return c
	
}

func receive(r <-chan int){
	for v := range r {
		fmt.Println(v)
	}
}