package main

import "fmt"

func main() {
	var numero int64 = -1000000000000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	//NUMEROS REAIS
	var numeroReal1 float32 = 123.32
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1222222222222.22
	fmt.Println(numeroReal2)

	//STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
}