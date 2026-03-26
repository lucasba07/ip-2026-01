package main

import "fmt"

func main() {
	var preco float64
	var codigo int
	var preoFinal float64

	fmt.Print("Digite o preço de etiqueta (R$): ")
	fmt.Scan(&preco)

	fmt.Print("Digite o código de pagamento:\n1 - À vista (dinheiro/cheque) - 10% desconto\n2 - À vista (cartão) - 5% desconto\n3 - 2 vezes - sem juros\n4 - 3 vezes - 10% de juros\nCódigo: ")
	fmt.Scan(&codigo)

	switch codigo {
	case 1:
		preoFinal = preco * 0.90  // 10% de desconto
		fmt.Printf("À vista (dinheiro/cheque) - 10%% de desconto\n")
		fmt.Printf("Preço original: R$ %.2f\n", preco)
		fmt.Printf("Preço final: R$ %.2f\n", preoFinal)

	case 2:
		preoFinal = preco * 0.95  // 5% de desconto
		fmt.Printf("À vista (cartão) - 5%% de desconto\n")
		fmt.Printf("Preço original: R$ %.2f\n", preco)
		fmt.Printf("Preço final: R$ %.2f\n", preoFinal)

	case 3:
		preoFinal = preco
		parcelado := preco / 2
		fmt.Printf("2 vezes - sem juros\n")
		fmt.Printf("Valor por parcela: R$ %.2f\n", parcelado)
		fmt.Printf("Total: R$ %.2f\n", preoFinal)

	case 4:
		preoFinal = preco * 1.10  // 10% de juros
		parcelado := preoFinal / 3
		fmt.Printf("3 vezes - 10%% de juros\n")
		fmt.Printf("Valor por parcela: R$ %.2f\n", parcelado)
		fmt.Printf("Total: R$ %.2f\n", preoFinal)

	default:
		fmt.Println("Erro: Código de pagamento inválido!")
	}
}
