package main

import (
	"github.com/jeffotoni/gocorreio.frete/config"
	"github.com/jeffotoni/gocorreio.frete/handlers"
	"log"
	"net/http"
)

func main() {
	//log.Println("version: 0.0.1")
	mux := http.NewServeMux()
	mux.HandleFunc("/frete", handler.Frete)
	mux.HandleFunc("/frete/", handler.NotFound)
	mux.HandleFunc("/", handler.NotFound)

	server := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	log.Println("Port:", config.Port, " Version:", config.VersionApp)
	log.Fatal(server.ListenAndServe())
}
