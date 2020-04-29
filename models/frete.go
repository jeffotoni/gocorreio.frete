package models

import "encoding/xml"

type ResultCServico struct {
	Codigo                string `json:"Codigo"`
	Valor                 string `json:"Valor"`
	PrazoEntrega          string `json:"PrazoEntrega"`
	ValorSemAdicionais    string `json:"ValorSemAdicionais"`
	ValorMaoPropria       string `json:"ValorMaoPropria"`
	ValorAvisoRecebimento string `json:"ValorAvisoRecebimento"`
	ValorValorDeclarado   string `json:"ValorValorDeclarado"`
	EntregaDomiciliar     string `json:"EntregaDomiciliar"`
	EntregaSabado         string `json:"EntregaSabado"`
	ObsFim                string `json:"obsFim"`
	Erro                  string `json:"Erro"`
	MsgErro               string `json:"MsgErro"`
	ValorTotal            string `json:"valorTotal"`
}

type GetFrete struct {
	NCdEmpresa          string  `json:"ncdempresa"`
	SDsSenha            string  `json:"sdssenha"`
	SCepOrigem          string  `json:"sceporigem"`
	SCepDestino         string  `json:"scepdestino"`
	NVlPeso             float32 `json:"nvlpeso"`
	NCdFormato          int     `json:"ncdformato"`
	NVlComprimento      int     `json:"nvlcomprimento"`
	NVlAltura           int     `json:"nvlaltura"`
	NVlLargura          int     `json:"nvllargura"`
	SCdMaoPropria       string  `json:"scdmaopropria"`
	NVlValorDeclarado   string  `json:"nvlvalordeclarado"`
	SCdAvisoRecebimento string  `json:"scdavisorecebimento"`
	//NCdServico          string   `json:"ncdservico"`
	NVlDiametro int      `json:"nvldiametro"`
	StrRetorno  string   `json:"strretorno"`
	Servicos    []string `json:"servicos"`
}

type ResultCServicoXML string

// type Result struct {
// 	Body []ResultCServicoMin
// }

type ServicosXML struct {
	XMLName  xml.Name `xml:"Servicos"`
	Text     string   `xml:",chardata"`
	CServico struct {
		Text                  string `xml:",chardata"`
		Codigo                string `xml:"Codigo"`
		Valor                 string `xml:"Valor"`
		PrazoEntrega          string `xml:"PrazoEntrega"`
		ValorSemAdicionais    string `xml:"ValorSemAdicionais"`
		ValorMaoPropria       string `xml:"ValorMaoPropria"`
		ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento"`
		ValorValorDeclarado   string `xml:"ValorValorDeclarado"`
		EntregaDomiciliar     string `xml:"EntregaDomiciliar"`
		EntregaSabado         string `xml:"EntregaSabado"`
		ObsFim                string `xml:"obsFim"`
		Erro                  string `xml:"Erro"`
		MsgErro               string `xml:"MsgErro"`
	} `xml:"cServico"`
}

var DefaultXmlError string = `<?xml version="1.0" encoding="ISO-8859-1" ?><Servicos><cServico><Codigo>%s</Codigo><Valor>0,00</Valor><PrazoEntrega>0</PrazoEntrega><ValorSemAdicionais>0,00</ValorSemAdicionais><ValorMaoPropria>0,00</ValorMaoPropria><ValorAvisoRecebimento>0,00</ValorAvisoRecebimento><ValorValorDeclarado>0,00</ValorValorDeclarado><EntregaDomiciliar></EntregaDomiciliar><EntregaSabado></EntregaSabado><obsFim></obsFim><Erro>%d</Erro><MsgErro><![CDATA[%s]]></MsgErro></cServico></Servicos>`
