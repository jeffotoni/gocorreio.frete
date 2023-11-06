package fretev2

import (
	"errors"

	"github.com/jeffotoni/gocorreio.frete/models"
)

func IsValid(gf *models.PostFretev2) error {
	if len(gf.ParametrosPrazo) == 0 {
		return errors.New(`{"msg":"ParametrosPrazo é obrigatório!"}`)
	}

	if len(gf.ParametrosProduto) == 0 {
		return errors.New(`{"msg":"ParametrosProduto é obrigatório!"}`)
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

	for _, params := range gf.ParametrosProduto {
		if len(params.CoProduto) <= 0 {
			return errors.New(`{"msg":"Campo CoProduto é obrigatório"}`)
		}
		if len(params.CepOrigem) <= 0 {
			return errors.New(`{"msg":"Campo CepOrigem é obrigatório"}`)
		}
		if len(params.CepDestino) <= 0 {
			return errors.New(`{"msg":"Campo CepDestino é obrigatório"}`)
		}
		if len(params.DtEvento) <= 0 {
			return errors.New(`{"msg":"Campo DtEvento é obrigatório"}`)
		}
		if len(params.PsObjeto) <= 0 {
			return errors.New(`{"msg":"Campo PsObjeto é obrigatório"}`)
		}
		if len(params.TpObjeto) <= 0 {
			return errors.New(`{"msg":"Campo TpObjeto é obrigatório"}`)
		}
		if params.Comprimento == 0 {
			return errors.New(`{"msg":"Campo Comprimento é obrigatório"}`)
		}
		if params.Largura == 0 {
			return errors.New(`{"msg":"Campo Largura é obrigatório"}`)
		}
		if params.Altura == 0 {
			return errors.New(`{"msg":"Campo Largura é obrigatório"}`)
		}
	}
	return nil
}
