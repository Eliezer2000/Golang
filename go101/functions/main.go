package main

import (
	"fmt"
	"strconv"
)

func hello(nome string) {
	fmt.Println("Hello ", nome)
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	total = a + i
	return
}
func main() {
	hello("Joel")
	fmt.Println(sum(10, 10))

	x, err := convertAndSum(10, "90")
	fmt.Println(x, err)
}