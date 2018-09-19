package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("cidades.csv")

	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo. Erro: ", err.Error())
		return
	}

	// Leitura do arquivo com Scanner
	// scanner := bufio.NewScanner(arquivo)

	// for scanner.Scan() {
	// 	linha := scanner.Text()
	// 	fmt.Println(linha)
	// }

	// Leitura do arquivo pelo recurso de 'csv'
	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()

	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com o leitor CSV. Erro: ", err.Error())
		return
	}

	for index, linha := range conteudo {
		fmt.Printf("Linha[%d] é: %s \n", index, linha)
	}

	arquivo.Close()
}
