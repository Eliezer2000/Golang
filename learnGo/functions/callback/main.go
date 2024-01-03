package main

import "fmt"

func main() {

	t := somentePares(soma, []int{1, 2, 33, 34, 56, 55, 90, 87, 86}...)
	fmt.Println(t)
	fmt.Println()
	i := somenteImpares(soma,  []int{1, 2, 33, 34, 56, 55, 90, 87, 86}...)
	fmt.Println(i)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y{
		if v % 2 != 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}