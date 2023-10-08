package main

import (
	"log"
	"net/http"

	"github.com/MrHenri/multithreadApi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{cep}", handlers.GetCep)
	log.Fatal(http.ListenAndServe(":8080", r))
}
