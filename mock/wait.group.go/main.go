package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var Servicos = []string{"01", "292", "39"}

type RS struct {
	Codigo string
	Valor  string
}

type Result struct {
	Body []RS
}

var r Result
var rs RS

func main() {

	var chResult = make(chan Result, len(Servicos))

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	wg.Add(len(Servicos))
	for _, servico := range Servicos {

		go func(wg *sync.WaitGroup, cancel context.CancelFunc, ctx context.Context, servico string, chResult chan<- Result) {
			NewRequestWithContextCorreioFrete(wg, ctx, cancel, servico, chResult)
		}(&wg, cancel, ctx, servico, chResult)
	}

	go func() {
		wg.Wait()
		c := <-chResult
		fmt.Println("done:", c)
	}()

	time.Sleep(time.Second * 11)
	//wg.Wait()
}

func NewRequestWithContextCorreioFrete(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, servico string, chResult chan<- Result) {
	defer wg.Done()
	fmt.Println("estou executando")

	rs.Codigo = servico
	rs.Valor = "100.22"
	r.Body = append(r.Body, rs)

	time.Sleep(time.Second * 5)
	chResult <- r

}
