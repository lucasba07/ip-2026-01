package main

import "fmt"

func isQuadradoPerfeito(n int) bool {
	for i := 0; i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

func main() {
	var(
		lista []int
		num int
	)

	for {
		fmt.Print("Informe um valor inteiro: ")
		fmt.Scanln(&num)

		if num <= 0 {
			fmt.Println("Encerrando programa...")
			return
		}

		lista = append(lista, num)

		if isQuadradoPerfeito(num) {
			fmt.Printf("O número %v é um quadrado perfeito\n", num)
		} else {
			fmt.Printf("O número %v não é um quadrado perfeito\n", num)
		}
	}
}