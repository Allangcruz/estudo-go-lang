package main

import "fmt"

func main() {

	// Criação de slices
	var numeros []int

	fmt.Println(numeros, len(numeros), cap(numeros))

	// Inicialzando e alterando valor do slice
	numeros = make([]int, 5)

	fmt.Println(numeros, len(numeros), cap(numeros))

	// Criando slice com inicialização
	capitais := []string{"Lisboa"}

	fmt.Println(capitais, len(capitais), cap(capitais))

	// Adicionando novo item na lista
	capitais = append(capitais, "Brasília")

	fmt.Println(capitais, len(capitais), cap(capitais))

	// Alterarndo um iten da lista
	capitais[1] = "Brasilia"

	fmt.Println(capitais, len(capitais), cap(capitais))

	//Cria slice com inicializações apos sua criacao
	cidades := make([]string, 4)
	cidades[0] = "New York"
	cidades[1] = "Londres"
	cidades[2] = "Tokio"
	cidades[3] = "Singapura"

	fmt.Println(cidades, len(cidades), cap(cidades))

	for index, cidade := range cidades {
		fmt.Printf("Cidade[%d] é %s\n\r", index, cidade)
	}

}
