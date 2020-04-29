package main

import (
	"fmt"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/frete"
)

func main() {
	var gf = &models.GetFrete{
		NCdEmpresa:          "codigo-empresa-aqui",
		SDsSenha:            "senha-empresa-aqui",
		SCepOrigem:          "01405001",
		SCepDestino:         "06765000",
		NVlPeso:             1.5,
		NCdFormato:          1,
		NVlComprimento:      28,
		NVlAltura:           4,
		NVlLargura:          13,
		SCdMaoPropria:       "N",
		NVlValorDeclarado:   "0,00",
		SCdAvisoRecebimento: "N",
		NVlDiametro:         0,
		StrRetorno:          "xml",
		Servicos:            []string{"04162", "04669", "1"},
	}

	result, err := frete.Search(gf)
	fmt.Println(err)
	fmt.Println(result)
}
