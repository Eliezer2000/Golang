package main

import "fmt"

func main() {

	total := soma()
	fmt.Println(total)
}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}