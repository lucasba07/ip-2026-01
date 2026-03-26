package main

import "fmt"

func main() {
	const salarioMinimo = 788.00
	const valorHoraExtra = 10.00

	var matricula int
	var horasExtras float64

	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scan(&matricula)

	fmt.Print("Digite a quantidade de horas-extras: ")
	fmt.Scan(&horasExtras)

	salarioHoraExtra := horasExtras * valorHoraExtra
	salarioBruto := (3 * salarioMinimo) + salarioHoraExtra

	var descontoINSS float64 = 0
	var descontoIR float64 = 0

	if salarioBruto > 1500.00 {
		descontoINSS = salarioBruto * 0.12
	}

	if salarioBruto > 2000.00 {
		descontoIR = salarioBruto * 0.20
	}

	deducoes := descontoINSS + descontoIR
	salarioLiquido := salarioBruto - deducoes

	fmt.Printf("\n--- Contra-cheque ---\n")
	fmt.Printf("Matrícula: %d\n", matricula)
	fmt.Printf("Horas-extras: %.2f\n", horasExtras)
	fmt.Printf("Salário hora-extra: R$ %.2f\n", salarioHoraExtra)
	fmt.Printf("Salário Bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Desconto INSS: R$ %.2f\n", descontoINSS)
	fmt.Printf("Desconto IR: R$ %.2f\n", descontoIR)
	fmt.Printf("Total de deduções: R$ %.2f\n", deducoes)
	fmt.Printf("Salário Líquido: R$ %.2f\n", salarioLiquido)
}
