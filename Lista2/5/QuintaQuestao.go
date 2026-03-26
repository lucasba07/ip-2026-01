package main
import ("fmt")

func main(){
	var numero int
	
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	if numero % 5 == 0 {
		fmt.Printf("O número %v é divisível por 5.\n", numero)
	} else {
		fmt.Printf("O número %v não é divisível por 5.\n", numero)
	}
}