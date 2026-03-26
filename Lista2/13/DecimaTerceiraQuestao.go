package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Digite um número de 3 casas: ")
	fmt.Scan(&numero)

	if numero < 100 || numero > 999 {
		fmt.Println("Erro: O número deve ter exatamente 3 casas!")
		return
	}

	// Extrai o algarismo da casa das dezenas
	dezenas := (numero / 10) % 10

	fmt.Printf("Algarismo da casa das dezenas: %d\n", dezenas)
}
