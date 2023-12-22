package main

import (
	"fmt"
	"maps"
)

func main() {

	//Para criar um mapa vazio, use o builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	//Defina pares chave/valor usando name[key] = val sintaxe típica.
	m["K"] = 7
	m["l"] = 8

	fmt.Println("map:", m)

	//Obtenha um valor para uma chave com name[key]
	v1 := m["k"]
	fmt.Println("v1:", v1)

	//Se a chave não existir, o valor zero do tipo de valor será retornado.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	//O builtin lenretorna o número de pares chave/valor quando chamado em um mapa.
	fmt.Println("len:", len(m))

	//O builtin deleteremove pares chave/valor de um mapa
	delete(m, "k2")
	fmt.Println("map:", m)

	//Para remover todos os pares chave/valor de um mapa, use o clearmétodo builtin.
	clear(m)
	fmt.Println("map:", m)

	/*
		O segundo valor de retorno opcional ao obter um valor de um mapa indica se a
		chave estava presente no mapa. Isso pode ser usado para eliminar a ambiguidade
		entre chaves ausentes e chaves com valores zero, como 0ou "".
		Aqui não precisávamos do valor em si, então o ignoramos com o identificador em
		branco _ .
	*/
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//Você também pode declarar e inicializar um novo mapa na mesma linha com esta sintaxe.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	//O mapspacote contém diversas funções utilitárias úteis para mapas
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
