package fretev2

import (
	"encoding/json"

	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/util"
)

func MergeResponse(resultPrazo models.RespPrazo200, resultPreco models.RespPreco200) (resStr string, err error) {
	var results = make(map[string]models.ResultCServico)
	var resultsResp []models.ResultCServico

	var msgErro string = ""
	var prazoEntrega string = ""
	var valor string = ""
	var valorTotal string = ""

	for _, prazoR := range resultPrazo {

		msgErro = ""
		prazoEntrega = ""

		// fmt.Println("prazoR.CoProduto ", prazoR.CoProduto, " - Entrega: ", prazoR.PrazoEntrega)

		if len(prazoR.TxErro) > 0 {
			msgErro = prazoR.TxErro
			prazoEntrega = ""
			// fmt.Println("msgErro ", msgErro)
		} else {
			prazoEntrega = util.Concat("", prazoR.PrazoEntrega)
		}

		results[prazoR.CoProduto] = models.ResultCServico{
			Codigo:       prazoR.CoProduto,
			PrazoEntrega: prazoEntrega,
			MsgErro:      msgErro,
		}
	}

	for _, precoR := range resultPreco {

		msgErro = ""
		valor = ""
		valorTotal = ""

		if len(results[precoR.CoProduto].MsgErro) > 0 {
			msgErro = results[precoR.CoProduto].MsgErro
		}

		if len(precoR.TxErro) > 0 {
			msgErro = util.Concat(msgErro, " - ", precoR.TxErro)
		}

		// fmt.Println("precoR.CoProduto ", precoR.CoProduto, " - Valor: R$ ", precoR.PcBase, " - Valor total: R$ ", precoR.PcFinal)

		if len(msgErro) > 0 {
			valor = "0,00"
			valorTotal = "0,00"
			// fmt.Println("msgErro ", msgErro)
		} else {
			valor = precoR.PcBase
			valorTotal = precoR.PcFinal
		}

		results[precoR.CoProduto] = models.ResultCServico{
			Codigo:                results[precoR.CoProduto].Codigo,
			PrazoEntrega:          results[precoR.CoProduto].PrazoEntrega,
			Valor:                 valor,
			ValorValorDeclarado:   "0,00",
			ValorTotal:            valorTotal,
			EntregaDomiciliar:     "S",
			EntregaSabado:         "N",
			ObsFim:                "",
			Erro:                  msgErro,
			MsgErro:               msgErro,
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
