package preco

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/jeffotoni/gocorreio.frete/config"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/token"
	"github.com/jeffotoni/gocorreio.frete/pkg/util"

	"time"
)

var urlReqPrecoPost string = `https://api.correios.com.br/preco/v1/nacional`

func PostPreco(gf *models.PostPreco) (resp200 models.RespPreco200, err error) {

	authToken, err := token.GetToken()
	if err != nil {
		return
	}

	payloadBuf := new(bytes.Buffer)
	err = json.NewEncoder(payloadBuf).Encode(&gf)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrecoPost, " - NewEncoder: ", err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.TimeoutSearchPreco)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", urlReqPrecoPost, payloadBuf)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrecoPost, " - NewRequestWithContext: ", err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", util.Concat("Bearer ", authToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrecoPost, " - DefaultClient: ", err.Error()))
		return
	}
	defer resp.Body.Close()

	bodyRes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrecoPost, " - io.ReadAll: ", err.Error()))
		return
	}
	// fmt.Println(urlReqPrecoPost, " - resp.StatusCode ", resp.StatusCode)

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		err = errors.New(util.Concat("url: ", urlReqPrecoPost, " - ERROR return api - body:", string(bodyRes)))
		return
	}

	// resp200 = string(bodyRes)
	err = json.Unmarshal(bodyRes, &resp200)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrecoPost, " - Unmarshal resp200 return api - body:", string(bodyRes)))
		return
	}
	return
}
