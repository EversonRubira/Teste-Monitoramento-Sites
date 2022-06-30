package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {
	exibeintroducao()
	leSitesDoArquivo()
	for {

		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs ...")
		case 0:
			fmt.Println("Saindo do Programa, Até mais ...")
			os.Exit(0)
		default:
			fmt.Println("Nao conheco esse comando")
			os.Exit(-1)
		}

	}

}

func exibeintroducao() {
	nome := "Everson"
	versao := 4.9
	fmt.Println("Hello World. Ola Sr.", nome)
	fmt.Println("Este programa tem a versao", versao)

	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do Programa")

	fmt.Println("")

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando ...")
	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site na posiçao", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("site", site, "Foi carregado com sucesso!")
	} else {
		fmt.Println("site", site, "esta com problemas, na execuçao. Status Code", resp.StatusCode)
	}

}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {

		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	fmt.Println(sites)

	arquivo.Close()

	return sites
}
