package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/jeffotoni/gocorreio.frete/pkg/frete"
)

var (
	Port = ":8086"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/frete/", HandlerFrete)
	mux.HandleFunc("/frete", NotFound)
	mux.HandleFunc("/", NotFound)

	server := &http.Server{
		Addr:    Port,
		Handler: mux,
	}

	log.Println("port", Port)
	log.Fatal(server.ListenAndServe())
}

func HandlerFrete(w http.ResponseWriter, r *http.Request) {

	freteStr := strings.Split(r.URL.Path[1:], "/")[1]
	if len(freteStr) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := frete.Search(freteStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
	return
}
