package main

import (
	"fmt"
	"math"
)

func main() {
	vetor := make([]float64, 15)

	for i := range vetor {
		var num int

		fmt.Print("Informe um número inteiro: ")
		fmt.Scanln(&num)

		if num < 0 {
			vetor[i] = -1
		} else {
			vetor[i] = math.Sqrt(float64(num))
		}
	}

	for _, v := range vetor {
		fmt.Printf("%.2f\n", v)
	}
}