package main

import (
	"fmt"
	"slices"
)

func main() {
	janela := make([]int, 24)
	corredor := make([]int, 24)

	for {
		if len(poltronasVazias(janela)) == 0 && len(poltronasVazias(corredor)) == 0 {
			fmt.Println("Não há poltronas disponíveis no ônibus!")
			break
		}

		escolha := 0
		fmt.Printf("\nQual é o local da poltrona que você deseja?\n")
		fmt.Println("1 - Janela | 2 - Corredor | 3 - Sair")
		fmt.Scanln(&escolha)

		poltrona := 0

		switch escolha {
		case 1:
			disponiveis := poltronasVazias(janela)
			if len(disponiveis) == 0 {
				fmt.Println("Não há poltronas disponíveis na janela!")
			} else {
				fmt.Printf("Poltronas disponíveis: %v\n", disponiveis)
				fmt.Print("Informe qual dessas poltronas você deseja comprar: ")
				fmt.Scanln(&poltrona)

				if slices.Contains(disponiveis, poltrona) {
					janela[poltrona-1] = 1
					fmt.Printf("Poltrona %v comprada!\n", poltrona)
				} else {
					fmt.Println("Essa poltrona não está disponível!")
				}
			}

		case 2:
			disponiveis := poltronasVazias(corredor)
			if len(disponiveis) == 0 {
				fmt.Println("Não há poltronas disponíveis no corredor!")
			} else {
				fmt.Printf("Poltronas disponíveis: %v\n", disponiveis)
				fmt.Print("Informe qual dessas poltronas você deseja comprar: ")
				fmt.Scanln(&poltrona)

				if slices.Contains(disponiveis, poltrona) {
					corredor[poltrona-1] = 1
					fmt.Printf("Poltrona %v comprada!\n", poltrona)
				} else {
					fmt.Println("Essa poltrona não está disponível!")
				}
			}

		case 3: 
			fmt.Println("Programa encerrado!")
			return

		default:
			fmt.Println("Número inválido!")
		}
	}
}

func poltronasVazias(local []int) []int {
	disponivel := make([]int, 0, 24)

	for i := range local {
		if local[i] == 0 {
			disponivel = append(disponivel, i+1)
		}
	}
	return disponivel
}