# Frete gocorreio.frete

Um simples pacote para buscar nos correios os fretes, onde você terá o custo do frete, o json que irá passar deverá conter a senha da sua conta do correio para que os valores correspondem a sua realidade.

Podendo implementar para ter uma saída ainda mais completa conforme sua necessidade, então fique a vontade em alterar conforme seu cenário.

O server é extremamente rápido, e usa cache em memória ele está configurado para 2G de Ram, caso queira alterar está tudo bonitinho no /config.

gocorreio.frete também poderá ser usado como Lib, ou seja você irá conseguir fazer um import em seu pkg/frete  e fazer a chamada direto do seu método em seu código.

## Usar como Lib
```go

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
      NVlPeso:             1.0,
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

```

Ou se preferir for criar seu próprio serviço e sua api basta fazer como exemplo abaixo:
Existe em examples dois exemplos de commo integrar a lib gocorreio.frete em seu projeto.

```bash
package main

import (
   "encoding/json"
   "github.com/jeffotoni/gocorreio.frete/models"
   "github.com/jeffotoni/gocorreio.frete/pkg/frete"
   "log"
   "net/http"
)

var (
   Port = ":8087"
)

func main() {

   mux := http.NewServeMux()
   mux.HandleFunc("/frete", HandlerFrete)
   mux.HandleFunc("/frete/", NotFound)
   mux.HandleFunc("/", NotFound)

   server := &http.Server{
      Addr:    Port,
      Handler: mux,
   }

   log.Println("port", Port)
   log.Fatal(server.ListenAndServe())
}

```

Você pode fazer seu próprio build usando Go, ou você poderá utilizar docker-compose. O server irá funcionar na porta 8087, mas caso queira alterar basta ir na pasta /config.

Para subir o serviço para seu Servidor ou sua máquina local basta compilar, e a porta 8087 será aberta para consumir o endpoint /api/v1/{etiqueta}

# Install gocorreio.frete

Caso queira utilizar ele como serviço, basta baixa-lo ou usar o docker para utilizado.

## linux bash
```bash
$ git clone https://github.com/jeffotoni/gocorreio.frete
$ cd gocorreio.frete
$ go build -ldflags="-s -w" 
$ ./gocorreio.frete
$ 2020/04/29 12:56:46 Port: :8087

```

## docker e docker-compose

Deixei um script para facilitar a criação de sua imagem, todos os arquivos estão na raiz, docker-compose.yaml, Dockerfile tudo que precisa para personalizar ainda mais se precisar.
Ao rodar o script ele irá fazer pull da imagem que encontra-se no hub.docker.
```bash

$ sh deploy.gocorreio.frete.sh

```

## Listando service
```bash
$ docker-compose ps
Creating gocorreio.frete ... done
Name    Command   State           Ports         
------------------------------------------------
gocorreio.frete   /gocorreio.frete    Up      0.0.0.0:8087->8087/tcp
-e Generated Run docker-compose [ok] 

```

## Executando sua API
```bash

$ curl -i \
-d '{
   "nCdEmpresa":"xxxxxxx",
   "sDsSenha":"xxxxxx",
   "sCepOrigem":"01405001",
   "sCepDestino":"06765000",
   "nVlPeso":1.0,
   "nCdFormato":1,
   "nVlComprimento":28,
   "nVlAltura":4,
   "nVlLargura":13,
   "sCdMaoPropria":"N",
   "nVlValorDeclarado":"0,00",
   "sCdAvisoRecebimento":"N",
   "nCdServico":"04162",
   "nVlDiametro":0,
   "StrRetorno":"xml",
   "servicos":["04162","04669","1"]
}'
-XPOST "http://localhost:8087/frete"
```

## Saida Json
```json

[
   {
      "Codigo":"1",
      "Valor":"0,00",
      "PrazoEntrega":"0",
      "ValorSemAdicionais":"0,00",
      "ValorMaoPropria":"0,00",
      "ValorAvisoRecebimento":"0,00",
      "ValorValorDeclarado":"0,00",
      "EntregaDomiciliar":"",
      "EntregaSabado":"",
      "obsFim":"",
      "Erro":"-888",
      "MsgErro":"Para este serviço só está disponível o cálculo do PRAZO.",
      "valorTotal":""
   },
   {
      "Codigo":"04162",
      "Valor":"10,98",
      "PrazoEntrega":"2",
      "ValorSemAdicionais":"10,98",
      "ValorMaoPropria":"0,00",
      "ValorAvisoRecebimento":"0,00",
      "ValorValorDeclarado":"0,00",
      "EntregaDomiciliar":"S",
      "EntregaSabado":"S",
      "obsFim":"",
      "Erro":"0",
      "MsgErro":"",
      "valorTotal":""
   },
   {
      "Codigo":"04669",
      "Valor":"0,00",
      "PrazoEntrega":"0",
      "ValorSemAdicionais":"0,00",
      "ValorMaoPropria":"0,00",
      "ValorAvisoRecebimento":"0,00",
      "ValorValorDeclarado":"0,00",
      "EntregaDomiciliar":"",
      "EntregaSabado":"",
      "obsFim":"",
      "Erro":"-888",
      "MsgErro":"Não foi encontrada precificação. ERP-007: CEP de origem nao pode postar para o CEP de destino informado(-1).",
      "valorTotal":""
   }
]

```

