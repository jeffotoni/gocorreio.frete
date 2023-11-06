package token

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jeffotoni/gocorreio.frete/config"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/util"

	"time"
)

var (
	urlReqTokenCartaoPostagemPost string = `https://api.correios.com.br/token/v1/autentica/cartaopostagem`

	USUARIO_MEU_CORREIOS = os.Getenv("USUARIO_MEU_CORREIOS")
	CODIGO_ACESSO        = os.Getenv("CODIGO_ACESSO")
	NUMERO_CARTAO        = os.Getenv("NUMERO_CARTAO")
)

func PostToken() (resp200 string, err error) {

	var gf models.PostToken
	gf.Numero = NUMERO_CARTAO

	authBasicToken := util.Concat(USUARIO_MEU_CORREIOS, ":", CODIGO_ACESSO)
	sEnc := base64.StdEncoding.EncodeToString([]byte(authBasicToken))

	payloadBuf := new(bytes.Buffer)
	err = json.NewEncoder(payloadBuf).Encode(&gf)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqTokenCartaoPostagemPost, " - NewEncoder: ", err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.TimeoutSearchToken)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", urlReqTokenCartaoPostagemPost, payloadBuf)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqTokenCartaoPostagemPost, " - NewRequestWithContext: ", err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", util.Concat("Basic ", sEnc))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqTokenCartaoPostagemPost, " - DefaultClient: ", err.Error()))
		return
	}
	defer resp.Body.Close()

	bodyRes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqTokenCartaoPostagemPost, " - io.ReadAll: ", err.Error()))
		return
	}
	fmt.Println(urlReqTokenCartaoPostagemPost, " - resp.StatusCode ", resp.StatusCode)

	if resp.StatusCode != 201 {
		err = errors.New(util.Concat("url: ", urlReqTokenCartaoPostagemPost, " statuscode:", resp.Status, " - body:", string(bodyRes)))
		return
	}

	resp200 = string(bodyRes)
	// err = json.Unmarshal(bodyRes, &resp200)
	// if err != nil {
	// 	err = errors.New(util.Concat("url: ", urlReqTokenCartaoPostagemPost, " - Unmarshal resp200 return api - body:", string(bodyRes)))
	// 	return
	// }
	return
}
