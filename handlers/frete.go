package handler

import (
	"encoding/json"
	"github.com/jeffotoni/gocorreio.frete/models"
	"github.com/jeffotoni/gocorreio.frete/pkg/frete"
	"net/http"
)

func Frete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	endpoint := r.URL.Path
	if endpoint != "/frete" {
		w.WriteHeader(http.StatusFound)
		return
	}

	var gf models.GetFrete
	err := json.NewDecoder(r.Body).Decode(&gf)
	if err != nil {
		http.Error(w, `{"msg":"Ocorreu um erro ao tentar decodificar o json recebido!"}`, http.StatusBadRequest)
		return
	}

	var StrRetorno string = "xml"

	if len(gf.NCdEmpresa) <= 0 {
		http.Error(w, "Campo nCdEmpresa é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(gf.SDsSenha) <= 0 {
		http.Error(w, "Campo sDsSenha é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(gf.SCepOrigem) <= 0 {
		http.Error(w, "Campo sCepOrigem é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(gf.SCepDestino) <= 0 {
		http.Error(w, "Campo sCepDestino é obrigatorio", http.StatusBadRequest)
		return
	}
	if gf.NVlPeso <= 0 {
		http.Error(w, "Campo nVlPeso é obrigatorio", http.StatusBadRequest)
		return
	}
	if gf.NCdFormato <= 0 {
		http.Error(w, "Campo nCdFormato é obrigatorio", http.StatusBadRequest)
		return
	}
	if gf.NVlComprimento <= 0 {
		http.Error(w, "Campo nVlComprimento é obrigatorio", http.StatusBadRequest)
		return
	}
	if gf.NVlAltura <= 0 {
		http.Error(w, "Campo nVlAltura é obrigatorio", http.StatusBadRequest)
		return
	}
	if gf.NVlLargura <= 0 {
		http.Error(w, "Campo nVlLargura é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(gf.SCdMaoPropria) <= 0 {
		http.Error(w, "Campo sCdMaoPropria é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(gf.NVlValorDeclarado) <= 0 {
		http.Error(w, "Campo nVlValorDeclarado é obrigatorio", http.StatusBadRequest)
		return
	}
	if len(gf.SCdAvisoRecebimento) <= 0 {
		http.Error(w, "Campo sCdAvisoRecebimento é obrigatorio", http.StatusBadRequest)
		return
	}
	// if len(gf.NCdServico) <= 0 {
	// 	http.Error(w, "Campo nCdServico é obrigatorio", http.StatusBadRequest)
	// 	return
	// }
	if gf.NVlDiametro < 0 {
		http.Error(w, "Campo nVlDiametro é obrigatorio", http.StatusBadRequest)
		return
	}

	if len(gf.Servicos) <= 0 {
		http.Error(w, "Campo Servicos é obrigatorio", http.StatusBadRequest)
		return
	}

	if len(gf.StrRetorno) <= 0 {
		gf.StrRetorno = StrRetorno
	}

	result, err := frete.Search(&gf)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
	return
}
