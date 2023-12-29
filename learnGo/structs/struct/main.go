package main

import "fmt"

type Cliente struct {
	name      string
	last_name string
	smoker    bool
}

func main() {

	cliente1 := Cliente{
		"Joel",
		"Santos",
		false,
	}
	cliente2 := Cliente{
		"Maria",
		"silva",
		true,
	}
	fmt.Println(cliente1)
	fmt.Println(cliente2)
}