package main
import ("fmt")

func main(){
	var numero int
	
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	if numero > 20 && numero < 90 {
		fmt.Printf("O número %v está compreendido entre 20 e 90.\n", numero)
	} else {
		fmt.Printf("O número %v não está compreendido entre 20 e 90.\n", numero)
	}
}
