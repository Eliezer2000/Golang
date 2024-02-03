package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Cria um novo roteador para lidar com solicitações HTTP.
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! This is a simple Golang web server.")
	})

	// Define um manipulador para a rota "/parametros", que lê os parâmetros da consulta da URL.
	mux.HandleFunc("/parametros",  func( w http.ResponseWriter, r *http.Request){
		// Obtém o valor do parâmetro "nome" da consulta da URL.
		nome := r.URL.Query().Get("nome")

		// Se o parâmetro "nome" não estiver presente, retorna uma mensagem de erro.
		if nome == "" {
			fmt.Fprintf(w, "Error: parameter 'name' not found in query.")
			return
		}

		// Formata uma mensagem de saudação usando o valor do parâmetro "nome".
		mensagem := fmt.Sprintf("Hello, %s! You have successfully sent the 'name' parameter.", nome)

		// Envia a mensagem de saudação para o cliente.
		fmt.Fprintf(w, mensagem)
	})

	// Inicia o servidor web na porta 8080.
	http.ListenAndServe(":8080", mux)
}

/*
	Requisição no Brownser
	http://localhost:8080/parametros?nome=Joel
*/
