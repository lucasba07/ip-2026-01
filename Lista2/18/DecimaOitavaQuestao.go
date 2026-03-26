package main

import "fmt"

func main() {
	var preco float64
	var categoria string
	var diaSemana int
	var precoFinal float64

	fmt.Print("Digite o preço normal do DVD (R$): ")
	fmt.Scan(&preco)

	fmt.Print("Digite a categoria (comum ou lançamento): ")
	fmt.Scan(&categoria)

	fmt.Print("Digite o dia da semana (1-Domingo, 2-Segunda, 3-Terça, 4-Quarta, 5-Quinta, 6-Sexta, 7-Sábado): ")
	fmt.Scan(&diaSemana)

	if diaSemana < 1 || diaSemana > 7 {
		fmt.Println("Erro: Dia da semana inválido!")
		return
	}

	precoFinal = preco

	// Aplicar desconto conforme o dia
	if diaSemana == 2 || diaSemana == 3 || diaSemana == 5 {
		precoFinal = preco * 0.60 // 40% de desconto
	}

	// Aplicar acréscimo se for lançamento
	if categoria == "lançamento" {
		precoFinal = precoFinal * 1.15 // 15% de acréscimo
	}

	desconto := ""
	if diaSemana == 2 || diaSemana == 3 || diaSemana == 5 {
		desconto = " (com 40% de desconto)"
	}

	adicional := ""
	if categoria == "lançamento" {
		adicional = " (com 15% de acréscimo)"
	}

	fmt.Printf("Preço original: R$ %.2f%s%s\n", preco, desconto, adicional)
	fmt.Printf("Preço final: R$ %.2f\n", precoFinal)
}
