package main

import "fmt"

func main() {
	var idade int

	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)

	var classificacao string

	if idade >= 0 && idade <= 2 {
		classificacao = "Recém-nascido"
	} else if idade >= 3 && idade <= 11 {
		classificacao = "Criança"
	} else if idade >= 12 && idade <= 19 {
		classificacao = "Adolescente"
	} else if idade >= 20 && idade <= 55 {
		classificacao = "Adulto"
	} else if idade > 55 {
		classificacao = "Idoso"
	} else {
		fmt.Println("Erro: Idade inválida!")
		return
	}

	fmt.Printf("Classificação: %s\n", classificacao)
}
