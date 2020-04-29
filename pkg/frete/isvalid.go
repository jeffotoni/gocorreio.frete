package frete

import (
	"errors"
	"github.com/jeffotoni/gocorreio.frete/models"
)

func IsValid(gf *models.GetFrete) error {
	if len(gf.NCdEmpresa) <= 0 {
		return errors.New(`{"msg":"Campo nCdEmpresa é obrigatorio"}`)

	}
	if len(gf.SDsSenha) <= 0 {
		return errors.New(`{"msg":"Campo sDsSenha é obrigatorio"}`)

	}
	if len(gf.SCepOrigem) <= 0 {
		return errors.New(`{"msg":"Campo sCepOrigem é obrigatorio"}`)

	}
	if len(gf.SCepDestino) <= 0 {
		return errors.New(`{"msg":"Campo sCepDestino é obrigatorio"}`)

	}
	if gf.NVlPeso <= float32(0) {
		return errors.New(`{"msg":"Campo nVlPeso é obrigatorio"}`)

	}
	if gf.NCdFormato <= 0 {
		return errors.New(`{"msg":"Campo nCdFormato é obrigatorio"}`)

	}
	if gf.NVlComprimento <= 0 {
		return errors.New(`{"msg":"Campo nVlComprimento é obrigatorio"}`)

	}
	if gf.NVlAltura <= 0 {
		return errors.New(`{"msg":"Campo nVlAltura é obrigatorio"}`)

	}
	if gf.NVlLargura <= 0 {
		return errors.New(`{"msg":"Campo nVlLargura é obrigatorio"}`)

	}
	if len(gf.SCdMaoPropria) <= 0 {
		return errors.New(`{"msg":"Campo sCdMaoPropria é obrigatorio"}`)

	}
	if len(gf.NVlValorDeclarado) <= 0 {
		return errors.New(`{"msg":"Campo nVlValorDeclarado é obrigatorio"}`)

	}
	if len(gf.SCdAvisoRecebimento) <= 0 {
		return errors.New(`{"msg":"Campo sCdAvisoRecebimento é obrigatorio"}`)

	}
	if gf.NVlDiametro < 0 {
		return errors.New(`{"msg":"Campo nVlDiametro é obrigatorio"}`)

	}
	if len(gf.Servicos) <= 0 {
		return errors.New(`{"msg":"Campo Servicos é obrigatorio"}`)

	}
	if len(gf.StrRetorno) <= 0 {
		gf.StrRetorno = "xml"
	}
	return nil
}
