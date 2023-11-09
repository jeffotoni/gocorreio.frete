package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/util"
)

func GetToken() (token string, err error) {
	fileName := "./credentials/credentials.json"
	fileExists := false

	if stat, err2 := os.Stat(fileName); err2 == nil && !stat.IsDir() {
		fileExists = true
	}

	var hasExpired bool
	var resp200 models.RespToken201

	if fileExists {
		file, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("ERROR (os.ReadFile) ", err)
			// return "", err
		}

		err = json.Unmarshal(file, &resp200)
		if err != nil {
			fmt.Println("ERROR (json.Unmarshal) ", err)
			// return "", err
		}

		if len(resp200.Token) > 0 {

			hasExpired, err = CheckTokenExpired(resp200.ExpiraEm)
			if err != nil {
				fmt.Println("err (json.Unmarshal) ", err)
				// return "", err
			}
		}
	}

	if !hasExpired && len(resp200.Token) > 0 {
		// Token nao expirou, retorna o token do arquivo
		token = resp200.Token
		return
	}

	bodyRes, err := PostToken()
	if err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(bodyRes), &resp200)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqTokenCartaoPostagemPost, " - Unmarshal resp200 return api - body:", string(bodyRes)))
		return "", err
	}

	token = resp200.Token

	data, err := json.Marshal(resp200)
	if err != nil {
		return
	}

	// Salva no arquivo o novo token retornado
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return
	}
	return
}
