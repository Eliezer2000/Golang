package main

import "fmt"

func main() {

	//The most basic type, with a single condition.
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("=====================================")
	//A classic initial/condition/after for loop.
	for i := 0; i <= 9; i++ {
		fmt.Println(i)
	}
	fmt.Println("=====================================")
	//for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("LOOP")
		break
	}
	fmt.Println("=====================================")
	//You can also continue to the next iteration of the loop.
	for x := 0; x <= 5; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}
}
