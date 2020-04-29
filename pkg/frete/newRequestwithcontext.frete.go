package frete

import (
	"context"
	"fmt"
	"github.com/jeffotoni/gocorreio.frete/models"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var endpoint string = `http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?nCdEmpresa=%s&sDsSenha=%s&sCepOrigem=%s&sCepDestino=%s&nVlPeso=%d&nCdFormato=%d&nVlComprimento=%d&nVlAltura=%d&nVlLargura=%d&sCdMaoPropria=%s&nVlValorDeclarado=%s&sCdAvisoRecebimento=%s&nCdServico=%s&nVlDiametro=%d&StrRetorno=%s`

func NewRequestWithContextCorreioFrete(wg *sync.WaitGroup, gf *models.GetFrete, nCdServico string, chResult chan<- string) {
	defer wg.Done()

	endpointNow := fmt.Sprintf(endpoint, gf.NCdEmpresa, gf.SDsSenha, gf.SCepOrigem,
		gf.SCepDestino, gf.NCdFormato, gf.NVlPeso, gf.NVlComprimento, gf.NVlAltura, gf.NVlLargura, gf.SCdMaoPropria, gf.NVlValorDeclarado,
		gf.SCdAvisoRecebimento, nCdServico, gf.NVlDiametro, gf.StrRetorno)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*7000)
	defer cancel()

	runtime.Gosched()
	req, err := http.NewRequestWithContext(ctx, "GET", endpointNow, nil)
	if err != nil {
		return
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		errXml := fmt.Sprintf(models.DefaultXmlError, nCdServico, 10, "Error, timout, url do correio nao respondeu.")
		chResult <- errXml
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errXml := fmt.Sprintf(models.DefaultXmlError, nCdServico, 11, "Error, nCdServico invalido")
		chResult <- errXml
		return
	}

	defer response.Body.Close()

	if len(string(body)) > 0 &&
		response.StatusCode == http.StatusOK {
		chResult <- string(body)
		return
	}

	errXml := fmt.Sprintf(models.DefaultXmlError, nCdServico, 13, "Error, nCdServico invalido")
	chResult <- errXml
	return
}