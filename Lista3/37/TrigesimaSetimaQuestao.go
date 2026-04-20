package main

import "fmt"

func main() {
	var num int

	for {
		fmt.Print("Informe um número octal qualquer: ")
		fmt.Scanln(&num)

		if num >= 0 {
			break
		}
		fmt.Println("INVÁLIDO!")
	}

	fmt.Print("Decimal: ")
	octalParaDecimal(num)
	fmt.Println()
}

func octalParaDecimal(num int) {
	soma := 0
	base := 1

	for num > 0 {
		digito := num % 10

		if digito > 7 {
			fmt.Print("Número inválido (não é octal)")
			return
		}

		soma += digito * base
		base *= 8
		num /= 10
	}

	fmt.Print(soma)
}