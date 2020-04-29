package main

import (
	"encoding/json"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/frete"
	"log"
	"net/http"
)

var (
	Port = ":8086"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/frete", HandlerFrete)
	mux.HandleFunc("/frete/", NotFound)
	mux.HandleFunc("/", NotFound)

	server := &http.Server{
		Addr:    Port,
		Handler: mux,
	}

	log.Println("port", Port)
	log.Fatal(server.ListenAndServe())
}

func HandlerFrete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	endpoint := r.URL.Path
	if endpoint != "/frete" {
		w.WriteHeader(http.StatusFound)
		return
	}

	var gf models.GetFrete
	err := json.NewDecoder(r.Body).Decode(&gf)
	if err != nil {
		http.Error(w, `{"msg":"Ocorreu um erro ao tentar decodificar o json recebido!"}`, http.StatusBadRequest)
		return
	}

	if err := frete.IsValid(&gf); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := frete.Search(&gf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	return
}
