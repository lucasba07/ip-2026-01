package main

import (
	"fmt"
	"math"
)

func main() {
	var numero float64

	fmt.Print("Digite um número qualquer: ")
	fmt.Scan(&numero)

	if numero < 0 {
		fmt.Printf("O número %v é negativo, portanto o seu quadrado é: %v.\n", numero, (numero * numero))
	} else {
		fmt.Printf("O número %v é positivo ou nulo, portanto a sua raiz quadrada é: %.2f.\n", numero, (math.Sqrt(numero)))
	}
}
