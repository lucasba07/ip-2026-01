package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64
	var volume, area float64

	fmt.Print("Escolha a figura:\n1 - Cone\n2 - Cilindro\n3 - Esfera\nOpção: ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1: // Cone
		fmt.Print("Digite o raio: ")
		fmt.Scan(&raio)
		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)

		volume = (math.Pi * raio * raio * altura) / 3
		area = math.Pi * raio * math.Sqrt(raio*raio+altura*altura)

		fmt.Printf("Cone:\n")
		fmt.Printf("Volume: %.2f\n", volume)
		fmt.Printf("Área da superfície: %.2f\n", area)

	case 2: // Cilindro
		fmt.Print("Digite o raio: ")
		fmt.Scan(&raio)
		fmt.Print("Digite a altura: ")
		fmt.Scan(&altura)

		volume = math.Pi * raio * raio * altura
		area = 2 * math.Pi * raio * altura

		fmt.Printf("Cilindro:\n")
		fmt.Printf("Volume: %.2f\n", volume)
		fmt.Printf("Área da superfície: %.2f\n", area)

	case 3: // Esfera
		fmt.Print("Digite o raio: ")
		fmt.Scan(&raio)

		volume = (4 / 3) * math.Pi * raio * raio * raio
		area = 4 * math.Pi * raio * raio

		fmt.Printf("Esfera:\n")
		fmt.Printf("Volume: %.2f\n", volume)
		fmt.Printf("Área da superfície: %.2f\n", area)

	default:
		fmt.Println("Erro: Opção inválida!")
	}
}
