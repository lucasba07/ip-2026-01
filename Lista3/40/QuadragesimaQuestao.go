package main

import "fmt"

func main() {
	ingressos := 130
	preco := 6.0

	var (
		lucroMax float64
		precoMax float64
		ingMax int
	)

	fmt.Println("PREÇO | INGRESSOS | LUCRO")
	fmt.Println("-------------------------------")

	for preco >= 1.0 {
		lucro := (float64(ingressos) * preco) - 300

		fmt.Printf("R$%.2f | %d | R$%.2f\n", preco, ingressos, lucro)

		if lucro > lucroMax || preco == 6.0 {
			lucroMax = lucro
			precoMax = preco
			ingMax = ingressos
		}

		preco -= 0.6
		ingressos += 30
	}

	fmt.Println("\n----- RESULTADO ------")
	fmt.Printf("Lucro máximo: R$%.2f\n", lucroMax)
	fmt.Printf("Preço ideal: R$%.2f\n", precoMax)
	fmt.Printf("Ingressos vendidos: %d\n", ingMax)
}