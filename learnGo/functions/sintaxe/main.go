package main

import "fmt"

func main() {

	//func (receiver) identifier(parameters) (returns) { code }

	valor := soma(10, 10)
	fmt.Println(valor)
	basica()
	n := "noite"
	argumento(n)


	
}


func basica(){
	fmt.Println("Bom dia!")
}

func argumento(s string){
	if s == "manha"{
		fmt.Println("Bom dia")
	} else if s == "tarde" {
		fmt.Println("Boa tarde")
	}else if s == "noite" {
		fmt.Println("Boa noite")
	} else {
		fmt.Println("tchau")
	}
}

func soma(x, y int) int {
	return x + y
}

/*
func som(x ...int) (int int) {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma , len(x)
}
*/