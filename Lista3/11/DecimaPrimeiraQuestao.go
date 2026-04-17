package main
import "fmt"

func main(){
	var num int

	for{
		fmt.Print("Informe um número inteiro positivo: ")
		fmt.Scanln(&num)
		if num < 0 {fmt.Println("VALOR INVÁLIDO!")} else{break}
	}
	res:=fatorial(num)
	fmt.Printf("%v\n", res)
}

func fatorial (num int) int {
	if num == 0 {
		return 1
	} 
	return num * fatorial(num-1)
}
