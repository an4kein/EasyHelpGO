package main

// Author: Eduardo Barbosa @anakein

import (
	"fmt"
	"strings"
)

var perguntas = []string{"", "1 - Problemas com wifi", "2 - Problemas com minha internet", "3 - Segunda Via da Fatura"}
var options int

func main() {
	fmt.Println("Enter your option: ")
	fmt.Println(strings.Join(perguntas, "\n"))
	fmt.Scanln(&options)
	meucanal := make(chan []string)
	go send(meucanal)
	receiver(meucanal)
}

func send(s chan<- []string) {
	for _, v := range perguntas {
		lm := []string{v}
		s <- lm
	}
}

func receiver(r <-chan []string) {
	if options == 1 {
		fmt.Println("Option 1: \nDesligue o seu roteador por 5 minutos")
	} else if options == 2 {
		fmt.Println("Option 2: \n Verifique se o roteador encontra-se ligado e com os cabos conectados")
	} else if options == 3 {
		fmt.Println("Option 3: \n Acesse o link http://example.com e imprima a segunda via")
	}
}
