package main

import "fmt"

func main() {
	var num int

	for{
		fmt.Print("Informe um número natural qualquer: ")
		fmt.Scanln(&num)

		if num > 0 {break}
		if num <= 0 {fmt.Println("INVÁLIDO!")}
	}

	fmt.Print("Binário: ")
	decimalParaBinario(num)
	fmt.Println()
}

func decimalParaBinario(num int) {
	if num == 0 {
		return
	}

	decimalParaBinario(num / 2)
	fmt.Print(num % 2)
}