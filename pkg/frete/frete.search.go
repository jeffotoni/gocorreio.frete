package frete

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/jeffotoni/gocorreio.frete/config"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/util"
	"github.com/jeffotoni/gocorreio.frete/service/ristretto"
)

func Search(gf *models.GetFrete) (string, error) {
	serviceOne := strings.Join(gf.Servicos, "")
	GSha1 := util.GSha1(util.Concat(serviceOne, "_", gf.NCdEmpresa, gf.SDsSenha, gf.SCepOrigem,
		gf.SCepDestino, gf.NCdFormato, gf.NVlPeso, gf.NVlComprimento, gf.NVlAltura, gf.NVlLargura, gf.SCdMaoPropria, gf.NVlValorDeclarado,
		gf.SCdAvisoRecebimento, gf.NVlDiametro, gf.StrRetorno))
	jsoncodigoFrete := ristretto.Get(GSha1)
	if len(jsoncodigoFrete) > 0 {
		//println("buscando em cache..")
		return jsoncodigoFrete, nil
	}

	runtime.GOMAXPROCS(config.NumCPU)
	var chResult = make(chan string, 1)

	var wg sync.WaitGroup
	for _, nCdServico := range gf.Servicos {
		wg.Add(1)
		go func(wg *sync.WaitGroup, gf *models.GetFrete, nCdServico string, chResult chan<- string) {
			NewRequestWithContextCorreioFrete(wg, gf, nCdServico, chResult)
		}(&wg, gf, nCdServico, chResult)
	}

	go func() {
		wg.Wait()
		close(chResult)
	}()

	var sjsonV []models.ResultCServico
	//var vazio, count int
	for t := range chResult {
		var sxml models.ServicosXML
		var sjson models.ResultCServico

		d := xml.NewDecoder(strings.NewReader(t))
		d.CharsetReader = util.CharsetReader
		err := d.Decode(&sxml)
		if err != nil {
			log.Println("Error NewDecoder: ", err.Error())
			continue
		}

		// if len(sxml.CServico.Valor) <= 0 ||
		// 	sxml.CServico.Valor == "0,00" ||
		// 	sxml.CServico.Valor == "0,0" {
		// 	vazio++
		// }

		sjson.Codigo = sxml.CServico.Codigo
		sjson.Valor = sxml.CServico.Valor
		sjson.PrazoEntrega = sxml.CServico.PrazoEntrega
		sjson.ValorSemAdicionais = sxml.CServico.ValorSemAdicionais
		sjson.ValorMaoPropria = sxml.CServico.ValorMaoPropria
		sjson.ValorAvisoRecebimento = sxml.CServico.ValorAvisoRecebimento
		sjson.ValorValorDeclarado = sxml.CServico.ValorValorDeclarado
		sjson.EntregaDomiciliar = sxml.CServico.EntregaDomiciliar
		sjson.EntregaSabado = sxml.CServico.EntregaSabado
		sjson.Erro = sxml.CServico.Erro
		sjson.MsgErro = sxml.CServico.MsgErro
		sjsonV = append(sjsonV, sjson)
		//count++
	}

	b, err := json.Marshal(sjsonV)
	if err != nil {
		return "", err
	}

	//println("vazio => ", vazio, " count:", count)
	jsoncodigoFrete = string(b)
	ristretto.SetTTL(GSha1, jsoncodigoFrete, time.Duration(config.TTLCacheFrete)*time.Second)

	return jsoncodigoFrete, nil
}
