# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 go build -ldflags="-s -w" -o gocorreio.frete main.go
	upx gocorreio.frete
	@echo "buid completo..."
	@echo "\033[0;33m################ run #####################\033[0m"
	./gocorreio.frete

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/gocorreio.frete
	CGO_ENABLED=0 go build -ldflags="-s -w" -o gocorreio.frete main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"
	@echo "fim"

deploy.docker:
	@echo "########## Compilando nossa API ... "
	sh deploy.docker.hub.sh
	@echo "fim"

deploy.aws:
	@echo "########## Compilando nossa API ... "
	#CGO_ENABLED=0 go build -ldflags="-s -w" -o gocorreio.frete main.go
	CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -ldflags="-s -w" -o gocorreio.frete main.go
	sh deploy.aws.sh
	@echo "fim"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
