package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MrHenri/multithreadApi/controllers"
	"github.com/gorilla/mux"
)

func GetCep(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	cep, err := controllers.ValidateCep(mux.Vars(r)["cep"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	go GetApiCep(controllers.CepInput{Cep: cep, Api: controllers.CDN}, ch)
	go GetApiCep(controllers.CepInput{Cep: cep, Api: controllers.VIA}, ch)

	select {
	case out := <-ch:
		fmt.Println(out)
	case <-time.After(time.Second):
		fmt.Print("timeout")
	}
}

func GetApiCep(cepInput controllers.CepInput, ch chan string) {
	req, err := cepInput.GetCepRequest()
	if err != nil {
		log.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	ch <- fmt.Sprintf("API: %s\nPayload:\n%s\n", cepInput.Api, string(body))
}
