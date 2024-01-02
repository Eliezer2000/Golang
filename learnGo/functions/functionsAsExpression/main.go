package main

import "fmt"

func main() {

	x := 300
	y := func(x int) int{
		//fmt.Println(x, " vezes 45 é : ")
		return x * 45
	}
	fmt.Println(x, "vezes 45 é: ", y(x))

}