package main

import "fmt"

func main() {
	var id int
	var nota1, nota2, nota3, mediaExercicios float64

	fmt.Print("Digite o número de identificação: ")
	fmt.Scan(&id)

	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&nota1)

	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&nota2)

	fmt.Print("Digite a terceira nota: ")
	fmt.Scan(&nota3)

	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scan(&mediaExercicios)

	mediaFinal := (nota1 + (nota2 * 2) + (nota3 * 3) + mediaExercicios) / 7

	var conceito string
	var status string

	if mediaFinal >= 9.0 && mediaFinal <= 10.0 {
		conceito = "A"
		status = "APROVADO"
	} else if mediaFinal >= 7.5 && mediaFinal < 9.0 {
		conceito = "B"
		status = "APROVADO"
	} else if mediaFinal >= 6.0 && mediaFinal < 7.5 {
		conceito = "C"
		status = "APROVADO"
	} else if mediaFinal >= 4.0 && mediaFinal < 6.0 {
		conceito = "D"
		status = "REPROVADO"
	} else {
		conceito = "E"
		status = "REPROVADO"
	}

	fmt.Printf("\n--- Resultado ---\n")
	fmt.Printf("Número do aluno: %d\n", id)
	fmt.Printf("Nota 1: %.2f\n", nota1)
	fmt.Printf("Nota 2: %.2f\n", nota2)
	fmt.Printf("Nota 3: %.2f\n", nota3)
	fmt.Printf("Média dos exercícios: %.2f\n", mediaExercicios)
	fmt.Printf("Média de aproveitamento: %.2f\n", mediaFinal)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Status: %s\n", status)
}
