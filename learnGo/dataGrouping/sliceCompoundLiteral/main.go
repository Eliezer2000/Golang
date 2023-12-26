package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice2 := append(slice, 8, 9)
	fmt.Println(slice2)

	fmt.Println(slice[2])
	slice[3] = 345786

	slice[20] = 1
	fmt.Println(slice[20]) 

}