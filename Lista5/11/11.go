package main

import (
	"fmt"
	"math"
)

func main(){
	vetor := make([]float64, 100)

	for i:= range vetor{
		fmt.Print("Informe um número qualquer: ")
		fmt.Scanln(&vetor[i])
	} 
	fmt.Printf("A soma resultou em: %.2f", somaVetor(vetor))
}

func somaVetor (v []float64) float64{
	cont := len(v) - 1
	soma:= 0.0
	for i:=0; i<50; i++ {
		soma += math.Pow((v[i] - v[cont]), 3)
		cont--
	}
	return soma
}
