package main

import "fmt"

/*
- Callback: passe uma função como argumento a outra função.

*/
func main() {
	x := somaPares(soma, []int{1, 2, 33, 34, 56, 55, 90, 87, 86}...)

	fmt.Println(x)

	PassarArgumento(argumento)

}

func argumento() {
	fmt.Println("Hello")
}

func PassarArgumento(x func()) {
	fmt.Println("World")
	x()
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	} 
	return n
}
func somaPares(f func(x ...int) int, y ...int) int{
	var slice []int
	for _, v := range y {
		if v % 2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}