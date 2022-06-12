package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 5
const delay = 5

func main() {
	introdução()
	for {
		Menu()
		comando := cmd()

		switch comando {
		case 0:
			fmt.Println("Saindo do programa! Sentirei sua falta ^^")
			os.Exit(0)
		case 1:
			verificação()
		case 2:
			fmt.Println("Exibindo Logs")
			exibirLogs()
		default:
			fmt.Println("Comando desconhecido!")
			os.Exit(-1)
		}
	}
}

func introdução() {

	nome := "Gustavo"
	versão := "2.0.1"
	fmt.Println("Olá, ", nome)
	fmt.Println("Este programa está atualmente na versão ", versão)
	fmt.Println("")
}

func Menu() {

	fmt.Println("0- Sair do Programa")
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("")
}

func cmd() int {

	var cmd int
	fmt.Scan(&cmd)
	fmt.Println("O comando escolhido foi", cmd)
	fmt.Println("")

	return cmd
}

func verificação() {
	fmt.Println("Monitoramento Inciado...")
	fmt.Println("")
	sites := lerArquivos()
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			print("Testando site: ", i, ": ", site)
			fmt.Println("")
			testarSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testarSite(site string) {

	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Um erro inesperado ocorreu! Erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "esta online!")
		registerLogs(site, true)
	} else {
		fmt.Println("O site", site, "não esta respondendo... adicionando codigo de resposta:", resp.StatusCode, "na log!")
		registerLogs(site, false)
	}
}

func lerArquivos() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Um erro inesperado ocorreu! Erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err != nil {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registerLogs(site string, status bool) {
	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 - 15:04:05 ") + site + " - Online " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func exibirLogs() {

	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
