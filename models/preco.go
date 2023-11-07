package models

type PostPreco struct {
	IDLote            string          `json:"idLote"`
	ParametrosProduto []ParamsProduto `json:"parametrosProduto"`
}

type ParamsProduto struct {
	CoProduto          string          `json:"coProduto"`
	NuRequisicao       int             `json:"nuRequisicao"`
	CepOrigem          string          `json:"cepOrigem"`
	PsObjeto           string          `json:"psObjeto"`
	TpObjeto           string          `json:"tpObjeto"`
	Comprimento        int             `json:"comprimento"`
	Largura            int             `json:"largura"`
	Altura             int             `json:"altura"`
	ServicosAdicionais []ServAdicional `json:"servicosAdicionais"`
	VlDeclarado        string          `json:"vlDeclarado"`
	DtEvento           string          `json:"dtEvento"`
	CepDestino         string          `json:"cepDestino"`
}

type ServAdicional struct {
	CoServAdicional string `json:"coServAdicional"`
}

type RespPreco200 []RespPreco

type RespPreco struct {
	CoProduto            string `json:"coProduto"`
	PcBase               string `json:"pcBase"`
	PcBaseGeral          string `json:"pcBaseGeral"`
	PeVariacao           string `json:"peVariacao"`
	PcReferencia         string `json:"pcReferencia"`
	VlBaseCalculoImposto string `json:"vlBaseCalculoImposto"`
	NuRequisicao         string `json:"nuRequisicao"`
	InPesoCubico         string `json:"inPesoCubico"`
	PsCobrado            string `json:"psCobrado"`
	PeAdValorem          string `json:"peAdValorem"`
	VlSeguroAutomatico   string `json:"vlSeguroAutomatico"`
	QtAdicional          string `json:"qtAdicional"`
	PcFaixa              string `json:"pcFaixa"`
	PcFaixaVariacao      string `json:"pcFaixaVariacao"`
	PcProduto            string `json:"pcProduto"`
	PcFinal              string `json:"pcFinal"`
	RespError
}
