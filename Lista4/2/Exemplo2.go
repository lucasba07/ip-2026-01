package main

import "fmt"

func somaRecursiva(lista []float64, i int) float64 {
	if i == len(lista) {
		return 0
	}
	return lista[i] + somaRecursiva(lista, i+1)
}

func main() {
	var tamanho int
	
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&tamanho)

	lista := make ([]float64, tamanho)

	for i:= range lista {
		fmt.Printf("Digite o valor para a posição %d: ", i)
		fmt.Scan(&lista[i])
	}

	soma := somaRecursiva(lista, 0)

	fmt.Printf("A soma dos valores do array é: %.2f\n", soma)
}