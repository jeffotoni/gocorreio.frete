package prazo

import (
	"errors"

	"github.com/jeffotoni/gocorreio.frete/models"
)

func IsValid(gf *models.PostPrazo) error {
	if len(gf.ParametrosPrazo) == 0 {
		return errors.New(`{"msg":"ParametrosPrazo é obrigatório!"}`)
	}

	for _, params := range gf.ParametrosPrazo {
		if len(params.CepDestino) <= 0 {
			return errors.New(`{"msg":"Campo CepDestino é obrigatório"}`)
		}
		if len(params.CepOrigem) <= 0 {
			return errors.New(`{"msg":"Campo CepOrigem é obrigatório"}`)
		}
		if len(params.CoProduto) <= 0 {
			return errors.New(`{"msg":"Campo CoProduto é obrigatório"}`)
		}
		if len(params.DtEvento) <= 0 {
			return errors.New(`{"msg":"Campo DtEvento é obrigatório"}`)
		}
	}
	return nil
}
