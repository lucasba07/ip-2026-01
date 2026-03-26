package main

import "fmt"

func main() {
	var preco float64
	var opcao string
	var total float64 = 0

	fmt.Print("Digite o preço inicial do carro: R$ ")
	fmt.Scan(&preco)

	total = preco

	for {
		fmt.Print("\nOpções:\na) Ar condicionado (R$ 1750,00)\nb) Pintura metálica (R$ 800,00)\nc) Vidro elétrico (R$ 1200,00)\nd) Direção hidráulica (R$ 2000,00)\ne) Finalizar\n\nEscolha uma opção: ")
		fmt.Scan(&opcao)

		switch opcao {
		case "a":
			total += 1750.00
			fmt.Println("Ar condicionado adicionado!")
		case "b":
			total += 800.00
			fmt.Println("Pintura metálica adicionada!")
		case "c":
			total += 1200.00
			fmt.Println("Vidro elétrico adicionado!")
		case "d":
			total += 2000.00
			fmt.Println("Direção hidráulica adicionada!")
		case "e":
			fmt.Printf("\nPreço final do carro: R$ %.2f\n", total)
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
