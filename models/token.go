package models

type PostToken struct {
	Numero string `json:"numero"`
}

type RespToken201 struct {
	Ambiente       string         `json:"ambiente"`
	ID             string         `json:"id"`
	IP             string         `json:"ip"`
	Perfil         string         `json:"perfil"`
	CNPJ           string         `json:"cnpj"`
	CartaoPostagem CartaoPostagem `json:"cartaoPostagem"`
	Emissao        string         `json:"emissao"`
	ExpiraEm       string         `json:"expiraEm"`
	ZoneOffset     string         `json:"zoneOffset"`
	Token          string         `json:"token"`
}

type CartaoPostagem struct {
	Numero   string `json:"numero"`
	Contrato string `json:"contrato"`
	Dr       int    `json:"dr"`
	Api      []int  `json:"api"`
}
