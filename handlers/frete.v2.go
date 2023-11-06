package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jeffotoni/gocorreio.frete/models"
	fretev2 "github.com/jeffotoni/gocorreio.frete/pkg/frete.v2"
)

func Fretev2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}
	endpoint := r.URL.Path
	if endpoint != "/v2/frete" {
		w.WriteHeader(http.StatusFound)
		return
	}

	var gf models.PostFretev2
	err := json.NewDecoder(r.Body).Decode(&gf)
	if err != nil {
		//println(err.Error())
		http.Error(w, `{"msg":"Ocorreu um erro ao tentar decodificar o json recebido!"}`, http.StatusBadRequest)
		return
	}

	if err := fretev2.IsValid(&gf); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// result, err := prazo.PostPrazo(&gf)
	// if err != nil {
	// 	// fmt.Println("ERROR: ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(result))
	// 	return
	// }

	// result, err := preco.PostPreco(&gf)
	// if err != nil {
	// 	// fmt.Println("ERROR: ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(result))
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	// w.Write([]byte(result))
	return
}
