package main

import (
	"fmt"
	"strings"
)

type Pessoa struct {
	nome string
	altura float64
	pesoIdeal float64
}

func main() {
	var pessoas []Pessoa

	for {
		var p Pessoa

		fmt.Printf("Informe o nome da pessoa: ")
		fmt.Scanln(&p.nome)

		if strings.EqualFold(p.nome, "FIM") {
			break
		}

		fmt.Printf("Informe a altura de %v (em metros): ", p.nome)
		fmt.Scanln(&p.altura)

		p.pesoIdeal = (72.7 * p.altura) - 58.0
		pessoas = append(pessoas, p)
	}

	fmt.Println("\nPessoas cadastradas:")
	for i, pessoa := range pessoas {
		fmt.Printf("%d: Nome=%s, Altura=%.2f, Peso Ideal=%.2f\n", i+1, pessoa.nome, pessoa.altura, pessoa.pesoIdeal)
	}
}
