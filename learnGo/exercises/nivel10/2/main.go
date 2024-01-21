package main

import "fmt"

/*
	- Faça esse código funcionar
*/

func main(){
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println("........\n")
	fmt.Printf("cs\t%T\n")

}