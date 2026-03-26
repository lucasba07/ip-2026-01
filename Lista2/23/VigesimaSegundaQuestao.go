package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)

	var classificacao string

	if idade < 16 {
		classificacao = "Não-eleitor"
	} else if idade >= 16 && idade < 18 {
		classificacao = "Eleitor facultativo"
	} else if idade >= 18 && idade <= 65 {
		classificacao = "Eleitor obrigatório"
	} else {
		classificacao = "Eleitor facultativo"
	}

	fmt.Printf("Classe eleitoral: %s\n", classificacao)
}
