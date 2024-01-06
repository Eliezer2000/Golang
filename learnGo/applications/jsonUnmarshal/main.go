package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Nome          string `json:"Nome"`
	Sobrenome     string `json:"Sobrenome"`
	Idade         int    `json:"Idade"`
	Profissao     string `json:"Profissao"`
	ContaBancaria int    `json:"ContaBancaria"`
}

func main() {
	sb := []byte(`{"Nome":"Joel","Sobrenome":"silva","Idade":55,"Profissao":"mecanico","ContaBancaria":10}`)
	var joel Info
	err := json.Unmarshal(sb, &joel)
	if err != nil {
		fmt.Println("Error :", err)
	}

	fmt.Println(joel)
	fmt.Println(joel.Profissao)
}