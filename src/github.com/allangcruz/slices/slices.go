package main

import "fmt"

func main() {

	// Criação de slices
	var numeros []int

	fmt.Println(numeros, len(numeros), cap(numeros))
	separador()

	// Inicialzando e alterando valor do slice
	numeros = make([]int, 5)

	fmt.Println(numeros, len(numeros), cap(numeros))
	separador()

	// Criando slice com inicialização
	capitais := []string{"Lisboa"}

	fmt.Println(capitais, len(capitais), cap(capitais))
	separador()

	// Adicionando novo item na lista
	capitais = append(capitais, "Brasília")

	fmt.Println(capitais, len(capitais), cap(capitais))
	separador()

	// Alterarndo um iten da lista
	capitais[1] = "Brasilia"

	fmt.Println(capitais, len(capitais), cap(capitais))
	separador()

	//Cria slice com inicializações apos sua criacao
	cidades := make([]string, 5)
	cidades[0] = "New York"
	cidades[1] = "Londres"
	cidades[2] = "Tokio"
	cidades[3] = "Madeira"
	cidades[4] = "Singapura"

	fmt.Println(cidades, len(cidades), cap(cidades))

	for index, cidade := range cidades {
		fmt.Printf("Cidade[%d] é %s\n\r", index, cidade)
	}
	separador()

	// -----------------------------------------------------------
	// SEGUNDA AULA

	// Recuperar o intervalo de um slice
	capitaisAsiaticas := cidades[3:5]

	fmt.Println(capitaisAsiaticas, len(capitaisAsiaticas), cap(capitaisAsiaticas))
	separador()

	apenasDoisIniciais := cidades[:2]

	fmt.Println(apenasDoisIniciais, len(apenasDoisIniciais), cap(apenasDoisIniciais))
	separador()

	apenasDoisFinais := cidades[len(cidades)-2:]

	fmt.Println(apenasDoisFinais, len(apenasDoisFinais), cap(apenasDoisFinais))
	separador()

	// Remover item do slice
	indexRemover := 2
	auxCidade := cidades[:indexRemover]
	auxCidade = append(auxCidade, cidades[indexRemover+1:]...)

	copy(cidades, auxCidade)

	fmt.Println(cidades, len(cidades), cap(cidades))
	separador()
}

// Exibe separador
func separador() {
	fmt.Println("-----------------------------------------------------")
}
