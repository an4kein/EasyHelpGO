package main

// Author: Eduardo Barbosa @anakein

import (
	"fmt"
	"strings"
)

var perguntas1 = []string{"Escolha uma opcao entre 1-3", "1 - Problemas com wifi", "2 - Problemas com minha internet", "3 - Segunda Via da Fatura"}
var options1 int
var respostas1 = []string{"Escolha uma opcao entre 1-3", "Option 1: \nDesligue o seu roteador por 5 minutos",
	"Option 2: \n Verifique se o roteador encontra-se ligado e com os cabos conectados",
	"Option 3: \n Acesse o link http://example.com e imprima a segunda via"}

func main() {
	fmt.Println("Enter your option: ")
	fmt.Println(strings.Join(perguntas1, "\n"))
	fmt.Scanln(&options1)
	meucanal := make(chan []string)
	go send1(meucanal)
	receiver1(meucanal)
}

func send1(s chan<- []string) {
	s <- respostas1
}

func receiver1(r <-chan []string) {
	for i, v := range <-r {
		if i == options1 {
			fmt.Println(v)
		}
	}
}
