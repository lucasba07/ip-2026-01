package main

import "fmt"

func main() {
	var n1, n2 int

	for {
		fmt.Print("Informe um número natural: ")
		fmt.Scanln(&n1)

		fmt.Print("Informe outro número natural: ")
		fmt.Scanln(&n2)

		if n1 > 0 && n2 > 0 {
			break
		}
		fmt.Println("VALORES INVÁLIDOS!")
	}

	fmt.Println("MMC:", mmc(n1, n2))
}

func mmc(a, b int) int {
	return (a * b) / mdc(a, b)
}

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}