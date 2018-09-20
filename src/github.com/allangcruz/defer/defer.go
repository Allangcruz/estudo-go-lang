package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/allangcruz/model"
)

func main() {
	arquivo, err := os.Open("cidades.csv")

	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo. Erro: ", err.Error())
		return
	}

	// Permite executar uma operção ao sair do programa
	defer arquivo.Close()

	// Leitura do arquivo pelo recurso de 'csv'
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()

	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com o leitor CSV. Erro: ", err.Error())
		return
	}

	// criar os arquivo
	arquivoJson, err := os.Create("cidades.json")

	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
		return
	}

	// Permite executar uma operção ao sair do programa
	defer arquivoJson.Close()

	escritor := bufio.NewWriter(arquivoJson)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for j, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]

			fmt.Printf("Cidade: %v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)

			if err != nil {
				fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
			}

			escritor.WriteString("	" + string(cidadeJSON))

			if (j + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}

	escritor.WriteString("\r\n]")
	escritor.Flush()
}
