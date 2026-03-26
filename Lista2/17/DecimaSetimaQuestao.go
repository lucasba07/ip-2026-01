package main

import "fmt"

func main() {
	var tipoConsumidor int
	var consumo float64
	var valor float64

	fmt.Print("Digite o tipo de consumidor (1-Residencial, 2-Comercial, 3-Industrial): ")
	fmt.Scan(&tipoConsumidor)

	fmt.Print("Digite o consumo de água (m³): ")
	fmt.Scan(&consumo)

	if consumo < 0 {
		fmt.Println("Erro: O consumo não pode ser negativo!")
		return
	}

	switch tipoConsumidor {
	case 1: // Residencial
		valor = 5.00 + (consumo * 0.05)
	case 2: // Comercial
		if consumo <= 80 {
			valor = 500.00
		} else {
			valor = 500.00 + ((consumo - 80) * 0.25)
		}
	case 3: // Industrial
		if consumo <= 100 {
			valor = 800.00
		} else {
			valor = 800.00 + ((consumo - 100) * 0.04)
		}
	default:
		fmt.Println("Erro: Tipo de consumidor inválido!")
		return
	}

	tipoNome := ""
	switch tipoConsumidor {
	case 1:
		tipoNome = "Residencial"
	case 2:
		tipoNome = "Comercial"
	case 3:
		tipoNome = "Industrial"
	}

	fmt.Printf("Tipo: %s\n", tipoNome)
	fmt.Printf("Consumo: %.2f m³\n", consumo)
	fmt.Printf("Valor a pagar: R$ %.2f\n", valor)
}
