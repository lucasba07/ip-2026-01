package main
import ("fmt")

func main(){
	var num1, num2 int
	
	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	soma := num1 + num2

	if soma > 20 {
		fmt.Printf("A soma dos números é maior que 20, portanto: %v.\n", (soma + 8))
	} else {
		fmt.Printf("A soma dos números é menor ou igual a 20, portanto: %v.\n", (soma - 5))
	}
}