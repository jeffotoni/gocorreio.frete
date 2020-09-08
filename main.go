package main

import (
	"log"
	"net/http"

	"github.com/jeffotoni/gcolor"
	"github.com/jeffotoni/gocorreio.frete/config"
	handler "github.com/jeffotoni/gocorreio.frete/handlers"
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

	log.Println(gcolor.YellowCor("Server Run Port"), config.Port, " Version:", config.VersionApp)
	log.Fatal(server.ListenAndServe())
}
