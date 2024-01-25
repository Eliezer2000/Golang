package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

	a, b := 10, 15
	switch {
	case a < b:
		fmt.Println("A é menor que B")
	case b > a:
		fmt.Println("B é maior que A") 
	default:
		fmt.Println("A e B são iguais!")
	}
}