package main	
import "fmt"

func main(){
	var num int
	for{
		fmt.Print("Informe um número inteiro maior ou igual a 3: ")
		fmt.Scanln(&num)
		if num < 3 {fmt.Println("VALOR INVÁLIDO!")} else{break}
	}
	sequencia:= make([]int, num)
	sequencia[0], sequencia[1] = 0, 1
	for i:= 2; i < num; i++{
		sequencia[i] = sequencia[i-1] + sequencia[i-2]
	}
	for i:= range sequencia{
		fmt.Printf("%v ",sequencia[i])
	}
}