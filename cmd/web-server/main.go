package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexarlindo/mini-web-server-go/pkg/handler"
)

func main() {
	//aqui vou add as rotas
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/hello", handler.HelloHandler)

	//vou fazer com que o server entregue paginas esteticas
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//iniciar o server
	port := 8000
	fmt.Printf("Server running na porta :  %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
