package fretev2

import (
	"encoding/json"

	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/util"
)

func MergeResponse(resultPrazo models.RespPrazo200, resultPreco models.RespPreco200) (resStr string, err error) {
	var results = make(map[string]models.ResultCServico)
	var resultsResp []models.ResultCServico

	for _, prazoR := range resultPrazo {
		// fmt.Println("prazoR.CoProduto ", prazoR.CoProduto, " - Entrega: ", prazoR.PrazoEntrega)
		results[prazoR.CoProduto] = models.ResultCServico{
			Codigo:       prazoR.CoProduto,
			PrazoEntrega: util.Concat("", prazoR.PrazoEntrega),
		}
	}

	for _, precoR := range resultPreco {
		// fmt.Println("precoR.CoProduto ", precoR.CoProduto, " - Valor: R$ ", precoR.PcBase, " - Valor total: R$ ", precoR.PcFinal)
		results[precoR.CoProduto] = models.ResultCServico{
			Codigo:                results[precoR.CoProduto].Codigo,
			PrazoEntrega:          results[precoR.CoProduto].PrazoEntrega,
			Valor:                 precoR.PcFinal,
			ValorValorDeclarado:   precoR.PcBase,
			ValorTotal:            precoR.PcFinal,
			EntregaDomiciliar:     "S",
			EntregaSabado:         "N",
			ObsFim:                "",
			Erro:                  "",
			MsgErro:               "",
			ValorSemAdicionais:    "0,00",
			ValorAvisoRecebimento: "0,00",
		}
	}

	for _, res := range results {
		resultsResp = append(resultsResp, res)
	}

	jResults, err := json.Marshal(resultsResp)
	if err != nil {
		// fmt.Println("ERROR: ", err)
		return
	}

	resStr = string(jResults)
	return
}
