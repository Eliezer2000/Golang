package main

import "fmt"

func main() {
	
	firstSlice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(firstSlice)

	firstSlice = append(firstSlice[:2], firstSlice[4:]...)

	fmt.Println(firstSlice)
}