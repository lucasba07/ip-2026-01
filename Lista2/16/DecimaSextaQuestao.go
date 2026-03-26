package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Digite o coeficiente A: ")
	fmt.Scan(&a)

	fmt.Print("Digite o coeficiente B: ")
	fmt.Scan(&b)

	fmt.Print("Digite o coeficiente C: ")
	fmt.Scan(&c)

	if a == 0 {
		fmt.Println("Erro: O coeficiente A não pode ser zero!")
		return
	}

	delta := (b * b) - (4 * a * c)

	fmt.Printf("Δ = %.2f\n", delta)

	if delta < 0 {
		fmt.Println("Classificação: RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		raiz := -b / (2 * a)
		fmt.Println("Classificação: RAIZ ÚNICA")
		fmt.Printf("x = %.2f\n", raiz)
	} else {
		raiz1 := (-b + math.Sqrt(delta)) / (2 * a)
		raiz2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("Classificação: RAÍZES DISTINTAS")
		fmt.Printf("x1 = %.2f\n", raiz1)
		fmt.Printf("x2 = %.2f\n", raiz2)
	}
}
