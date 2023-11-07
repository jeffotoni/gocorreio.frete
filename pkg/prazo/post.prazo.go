package prazo

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

var urlReqPrazoPost string = `https://api.correios.com.br/prazo/v1/nacional`

func PostPrazo(gf *models.PostPrazo) (resp200 models.RespPrazo200, err error) {

	authToken, err := token.GetToken()
	if err != nil {
		return
	}

	payloadBuf := new(bytes.Buffer)
	err = json.NewEncoder(payloadBuf).Encode(&gf)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrazoPost, " - NewEncoder: ", err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.TimeoutSearchPrazo)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", urlReqPrazoPost, payloadBuf)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrazoPost, " - NewRequestWithContext: ", err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", util.Concat("Bearer ", authToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrazoPost, " - DefaultClient: ", err.Error()))
		return
	}
	defer resp.Body.Close()

	bodyRes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrazoPost, " - io.ReadAll: ", err.Error()))
		return
	}
	// fmt.Println(urlReqPrazoPost, " - resp.StatusCode ", resp.StatusCode)

	if resp.StatusCode != 200 {
		err = errors.New(util.Concat("url: ", urlReqPrazoPost, " statuscode:", resp.Status, " - body:", string(bodyRes)))
		return
	}

	// resp200 = string(bodyRes)
	err = json.Unmarshal(bodyRes, &resp200)
	if err != nil {
		err = errors.New(util.Concat("url: ", urlReqPrazoPost, " - Unmarshal resp200 return api - body:", string(bodyRes)))
		return
	}
	return
}
