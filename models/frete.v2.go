package models

type PostFretev2 struct {
	IDLote            string          `json:"idLote"`
	ParametrosPrazo   []ParamsPrazo   `json:"parametrosPrazo"`
	ParametrosProduto []ParamsProduto `json:"parametrosProduto"`
}
