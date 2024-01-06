package main

import "fmt"

func main() {
	x := 10
	fmt.Println(&x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Printf("%T , %T ", x, y)
}