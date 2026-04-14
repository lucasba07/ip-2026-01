package main

import "fmt"

func main() {
	var (
		base float64
		expoente int
	)

	for {
		fmt.Print("Informe a base: ")
		fmt.Scanln(&base)
		fmt.Print("Informe o expoente: ")
		fmt.Scanln(&expoente)
		
		if base != 0 && expoente >= 0 {
			resultado := potencia(base, expoente)
			fmt.Printf("O resultado de %.2f elevado a %d é %.2f\n", base, expoente, resultado)
			break
		} else {
			fmt.Println("Base deve ser diferente de zero e expoente deve ser um número inteiro não negativo. Tente novamente.")
		}
	}
}

func potencia(base float64, expoente int) float64 {
	if expoente == 0 {
		return 1
	}
	return base * potencia(base, expoente-1)
}