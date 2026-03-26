package main

import "fmt"

func main() {
	var x float64

	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)

	if x == 2 {
		fmt.Println("Erro: x não pode ser igual a 2 (divisão por zero)!")
		return
	}

	fx := 8 / (2 - x)

	fmt.Printf("f(x) = 8 / (2 - x)\n")
	fmt.Printf("f(%.2f) = %.2f\n", x, fx)
}
