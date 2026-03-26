package main
import ("fmt")

func main()	{
	var num1, num2, num3, menor, maior, inter int
	
	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)
	fmt.Print("Digite o terceiro número inteiro: ")
	fmt.Scan(&num3)

	switch {
	case num1 < num2 && num1 < num3:
		menor = num1
		if num2 < num3 {
			inter = num2
			maior = num3
		} else {
			inter = num3
			maior = num2
		}
	case num2 < num1 && num2 < num3:
		menor = num2
		if num1 < num3 {
			inter = num1
			maior = num3
		} else {
			inter = num3
			maior = num1
		}
	default:
		menor = num3
		if num1 < num2 {
			inter = num1
			maior = num2
		} else {
			inter = num2
			maior = num1
		}
	}

	fmt.Printf("Os números em ordem crescente são: %v, %v e %v.\n", menor, inter, maior)
}