package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		showError(err, "main")
		return
	}

	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			showError(err, "main")
			return
		}
		fmt.Println(string(corpo))

	}
}

func showError(err error, scope string) {
	fmt.Println("[", scope, "] - Error identificado: ", err.Error())
}
