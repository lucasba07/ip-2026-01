package main

import "fmt"

func main() {
	var dia, mes, ano int

	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)

	fmt.Print("Digite o mês (1-12): ")
	fmt.Scan(&mes)

	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)

	if dia < 1 || dia > 31 || mes < 1 || mes > 12 {
		fmt.Println("Erro: Data inválida!")
		return
	}

	meses := [13]string{"", "janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}

	fmt.Printf("Data: %d de %s de %d\n", dia, meses[mes], ano)
}
