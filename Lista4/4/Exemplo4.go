package main

import "fmt"

func main() {
	var num int

	fmt.Print("Informe um número decimal qualquer: ")
	fmt.Scanln(&num)

	if num == 0 {
		fmt.Println("Binário: 0")
		return
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