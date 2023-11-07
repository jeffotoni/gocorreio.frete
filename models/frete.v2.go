package models

type PostFretev2 struct {
	IDLote            string          `json:"idLote"`
	ParametrosPrazo   []ParamsPrazo   `json:"parametrosPrazo"`
	ParametrosProduto []ParamsProduto `json:"parametrosProduto"`
}

type RespError struct {
	Msgs   []string `json:"msgs"`
	Date   string   `json:"date"`
	Method string   `json:"method"`
	Path   string   `json:"path"`
	TxErro string   `json:"txErro"`
}
