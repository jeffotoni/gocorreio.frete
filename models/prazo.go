package models

type PostPrazo struct {
	IDLote          string        `json:"idLote"`
	ParametrosPrazo []ParamsPrazo `json:"parametrosPrazo"`
}

type ParamsPrazo struct {
	CepDestino   string `json:"cepDestino"`
	CepOrigem    string `json:"cepOrigem"`
	CoProduto    string `json:"coProduto"`
	NuRequisicao string `json:"nuRequisicao"`
	DtEvento     string `json:"dtEvento"`
}

type RespPrazo200 []RespPrazo

type RespPrazo struct {
	CoProduto         string `json:"coProduto"`
	NuRequisicao      string `json:"nuRequisicao"`
	PrazoEntrega      int    `json:"prazoEntrega"`
	DataMaxima        string `json:"dataMaxima"`
	EntregaDomiciliar string `json:"entregaDomiciliar"`
	EntregaSabado     string `json:"entregaSabado"`
	RespError
}
