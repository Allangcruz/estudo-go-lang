package main

import "fmt"

func main() {

	// Criação de array definindo o tamanho
	var numeros [3]int

	numeros[0] = 3
	numeros[1] = 2
	numeros[2] = 1

	fmt.Println("\nQual a capacidade deste array: ", len(numeros))

	// Criação de array definir o tamanho e atribuição
	cantores := [2]string{"Michael Jackson", "John Lenom"}

	fmt.Printf("\nO que há nesse Array? \n%v\r\n", cantores)

	// Criação de array sem definir o tamanho e atribuição
	capitais := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}

	fmt.Printf("\nQual a capacidade deste array? \n%v\r\n\n", capitais)

	for index, capital := range capitais {
		fmt.Printf("Capital[%d] é %s\n\r", index, capital)
	}
}
