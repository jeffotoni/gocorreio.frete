package handler

import (
	"encoding/json"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/frete"
	"io/ioutil"
	"net/http"
	"strings"
)

func Frete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	endpoint := r.URL.Path
	if endpoint != "/frete" {
		w.WriteHeader(http.StatusFound)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		println("ReadAll: ", err.Error())
		http.Error(w, `{"msg":"Ocorreu um erro no ReadAll"}`, http.StatusBadRequest)
		return
	}
	println("..........................")
	println(string(b))

	var gf models.GetFrete
	err = json.NewDecoder(strings.NewReader(string(b))).Decode(&gf)
	if err != nil {
		println()
		println(err.Error())
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
