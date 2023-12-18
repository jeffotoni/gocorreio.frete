package util

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jeffotoni/gocorreio.frete/models"
)

func PrintErrorFretev2(msg string, gf *models.PostFretev2) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z")

	reqBody, err := json.Marshal(&gf)
	if err != nil {
		fmt.Println(timestamp, " - ERROR: ", msg, " - ", err.Error())
		return
	}

	fmt.Println(timestamp, " - ERROR: ", msg, " - JSON: ", string(reqBody))
	return
}

func PrintReqRespFretev2(gf *models.PostFretev2, resultPrazo models.RespPrazo200, resultPreco models.RespPreco200) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z")

	reqBody, err := json.Marshal(&gf)
	if err != nil {
		fmt.Println(timestamp, " - ERROR: ", err.Error())
		return
	}

	respPrazo, err := json.Marshal(&resultPrazo)
	if err != nil {
		fmt.Println(timestamp, " - ERROR: ", err.Error())
		return
	}

	respPreco, err := json.Marshal(&resultPreco)
	if err != nil {
		fmt.Println(timestamp, " - ERROR: ", err.Error())
		return
	}

	fmt.Println("=============================================")
	fmt.Println(timestamp, " - INFO - REQUEST BODY: ", string(reqBody))
	fmt.Println(timestamp, " - RESPONSE PRAZO: ", string(respPrazo))
	fmt.Println(timestamp, " - RESPONSE PRECO: ", string(respPreco))
	fmt.Println("=============================================")
	return
}
