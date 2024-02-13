package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}


func main() {
	generica("num")
	generica(123)
	generica(true)
}