package main
import ("fmt")

func main(){
	var num1, num2 int
	
	fmt.Print("Digite o primeiro número inteiro: ")
	fmt.Scan(&num1)
	fmt.Print("Digite o segundo número inteiro: ")
	fmt.Scan(&num2)

	if num1 % num2 == 0 {
		fmt.Printf("O número %v é divisível por %v.\n", num1, num2)
	} else if num2 % num1 == 0 {
		fmt.Printf("O número %v é divisível por %v.\n", num2, num1)
	} else {
		fmt.Printf("Nenhum dos números é divisível pelo outro.\n")
	}

}
