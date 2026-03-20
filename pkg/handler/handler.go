package handler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/static/index.html")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		nome = "visitante"
	}
	fmt.Fprintf(w, "Ola, %s! Bem-vindo ao servidor Go.", nome)

}
