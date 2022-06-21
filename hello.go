package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeintroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs ...")
	case 0:
		fmt.Println("Saindo do Programa ...")
		os.Exit(0)
	default:
		fmt.Println("Nao conheco esse comando")
		os.Exit(-1)
	}

}

func exibeintroducao() {
	nome := "Everson"
	versao := 4.9
	fmt.Println("Hello World. Ola Sr.", nome)
	fmt.Println("Este programa tem a versao", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")
	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)

}
