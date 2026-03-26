package main

import "fmt"

func main() {
	var valorCompra float64

	fmt.Print("Digite o valor da compra: R$ ")
	fmt.Scan(&valorCompra)

	if valorCompra <= 0 {
		fmt.Println("Erro: O valor da compra deve ser positivo!")
		return
	}

	var lucro float64
	if valorCompra < 10.00 {
		lucro = 0.70
	} else if valorCompra < 30.00 {
		lucro = 0.50
	} else if valorCompra < 50.00 {
		lucro = 0.40
	} else {
		lucro = 0.30
	}

	valorVenda := valorCompra * (1 + lucro)

	fmt.Printf("Valor da compra: R$ %.2f\n", valorCompra)
	fmt.Printf("Lucro: %.0f%%\n", (lucro * 100))
	fmt.Printf("Valor da venda: R$ %.2f\n", valorVenda)
}
