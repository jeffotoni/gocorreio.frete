package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jeffotoni/gocorreio.frete/models"
	fretev2 "github.com/jeffotoni/gocorreio.frete/pkg/frete.v2"
	"github.com/jeffotoni/gocorreio.frete/pkg/prazo"
	"github.com/jeffotoni/gocorreio.frete/pkg/preco"
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

	// var results map[string]models.ResultCServico

	var gfPrazo models.PostPrazo
	gfPrazo.IDLote = gf.IDLote
	gfPrazo.ParametrosPrazo = gf.ParametrosPrazo

	resultPrazo, err := prazo.PostPrazo(&gfPrazo)
	if err != nil {
		// fmt.Println("ERROR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var respPrazo200 models.RespPrazo200
	err = json.Unmarshal([]byte(resultPrazo), &respPrazo200)
	if err != nil {
		// fmt.Println("ERROR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// for _, prazoR := range respPrazo200 {
	// 	results[prazoR.CoProduto] = models.ResultCServico{
	// 		Codigo:       prazoR.CoProduto,
	// 		PrazoEntrega: util.Concat("", prazoR.PrazoEntrega),
	// 	}
	// }

	// fmt.Println("Result (prazo): ", resultPrazo)

	var gfPreco models.PostPreco
	gfPreco.IDLote = gf.IDLote
	gfPreco.ParametrosProduto = gf.ParametrosProduto

	resultPreco, err := preco.PostPreco(&gfPreco)
	if err != nil {
		// fmt.Println("ERROR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var respPreco200 models.RespPreco200
	err = json.Unmarshal([]byte(resultPreco), &respPreco200)
	if err != nil {
		// fmt.Println("ERROR: ", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// for _, precoR := range respPreco200 {
	// 	results[precoR.CoProduto] = models.ResultCServico{
	// 		Codigo:                precoR.CoProduto,
	// 		Valor:                 precoR.PcBase,
	// 		ValorValorDeclarado:   precoR.PcBase,
	// 		ValorTotal:            precoR.PcFinal,
	// 		EntregaDomiciliar:     "S",
	// 		EntregaSabado:         "N",
	// 		ObsFim:                "",
	// 		Erro:                  "",
	// 		MsgErro:               "",
	// 		ValorSemAdicionais:    "0,00",
	// 		ValorAvisoRecebimento: "0,00",
	// 	}
	// }

	// fmt.Println("Result (preco): ", resultPreco)

	// jResults, err := json.Marshal(results)
	// if err != nil {
	// 	// fmt.Println("ERROR: ", err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	// w.Write(jResults)
	return
}
