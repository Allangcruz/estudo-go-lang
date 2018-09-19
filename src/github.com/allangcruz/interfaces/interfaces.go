package main

import (
	"fmt"

	"github.com/allangcruz/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "jojo da Silva"

	queroAcordarComUmCacarejo(jojo)

	queroOuvirUmaPatNoLagoa(jojo)
}

func queroAcordarComUmCacarejo(galinha model.Galinha) {
	fmt.Println(galinha.Cacareja())
}

func queroOuvirUmaPatNoLagoa(pato model.Pato) {
	fmt.Println(pato.Quack())
}
