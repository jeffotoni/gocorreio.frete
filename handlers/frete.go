package handler

import (
	"encoding/json"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/frete"
	"net/http"
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

	var gf models.GetFrete
	err := json.NewDecoder(r.Body).Decode(&gf)
	if err != nil {
		//println(err.Error())
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
