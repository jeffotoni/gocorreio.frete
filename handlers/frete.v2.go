package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/jeffotoni/gocorreio.frete/models"
	fretev2 "github.com/jeffotoni/gocorreio.frete/pkg/frete.v2"
	"github.com/jeffotoni/gocorreio.frete/pkg/prazo"
	"github.com/jeffotoni/gocorreio.frete/pkg/preco"
	"github.com/jeffotoni/gocorreio.frete/pkg/util"
)

var (
	DISABLE_APP_LOG = os.Getenv("DISABLE_APP_LOG")
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
		util.PrintErrorFretev2(util.Concat("handlers/frete.v2.go - ", err.Error()), &gf)
		http.Error(w, `{"msg":"Ocorreu um erro ao tentar decodificar o json recebido!"}`, http.StatusBadRequest)
		return
	}

	if err := fretev2.IsValid(&gf); err != nil {
		util.PrintErrorFretev2(util.Concat("handlers/frete.v2.go - IsValid: ", err.Error()), &gf)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var gfPrazo models.PostPrazo
	gfPrazo.IDLote = gf.IDLote
	gfPrazo.ParametrosPrazo = gf.ParametrosPrazo

	resultPrazo, err := prazo.PostPrazo(&gfPrazo)
	if err != nil {
		util.PrintErrorFretev2(util.Concat("handlers/frete.v2.go - PostPrazo: ", err.Error()), &gf)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var gfPreco models.PostPreco
	gfPreco.IDLote = gf.IDLote
	gfPreco.ParametrosProduto = gf.ParametrosProduto

	resultPreco, err := preco.PostPreco(&gfPreco)
	if err != nil {
		util.PrintErrorFretev2(util.Concat("handlers/frete.v2.go - PostPreco: ", err.Error()), &gf)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	results, err := fretev2.MergeResponse(resultPrazo, resultPreco)
	if err != nil {
		util.PrintErrorFretev2(util.Concat("handlers/frete.v2.go - MergeResponse: ", err.Error()), &gf)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if DISABLE_APP_LOG == "true" {
		util.PrintReqRespFretev2(&gf, resultPrazo, resultPreco)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(results))
	return
}
