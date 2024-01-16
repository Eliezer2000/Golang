package main

import "fmt"

/*
	- Faça esse código funcionar:
    - Usando uma função anônima auto-executável
    - Usando buffer
*/
func main() {
	c := make(chan int)
	go func ()  {
		c <- 42			//- Usando uma função anônima auto-executável
	}()
	fmt.Println(<-c)


	d := make(chan int, 1)
	d <- 33				// - Usando buffer
	fmt.Println(d)

}